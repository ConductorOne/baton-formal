package connector

import (
	"context"
	"errors"
	"fmt"

	corev1 "buf.build/gen/go/formal/core/protocolbuffers/go/core/v1"
	"connectrpc.com/connect"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	"github.com/conductorone/baton-sdk/pkg/types/grant"
	"github.com/formalco/go-sdk/sdk/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

const PageSize = 100

type groupBuilder struct {
	client *sdk.FormalSDK
}

func (o *groupBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return groupResourceType
}

func (o *groupBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	request := connect.NewRequest(&corev1.ListGroupsRequest{
		Limit:  PageSize,
		Cursor: pToken.Token,
	})

	response, err := o.client.GroupServiceClient.ListGroups(ctx, request)
	if err != nil {
		return nil, "", nil, fmt.Errorf("GroupServiceClient.ListGroups error: %w", err)
	}

	var groups []*v2.Resource
	for _, group := range response.Msg.Groups {
		resourceGroup, err := formalGroupToResourceGroup(parentResourceID, group)
		if err != nil {
			return nil, "", nil, fmt.Errorf("formalGroupToResourceGroup error: %w", err)
		}
		groups = append(groups, resourceGroup)
	}

	token := ""
	if len(response.Msg.Groups) > 0 {
		token = response.Msg.Groups[len(response.Msg.Groups)-1].Id
	}

	rateLimit, err := rateLimitAnnotations(response.Header())
	if err != nil {
		return nil, "", nil, fmt.Errorf("rateLimitAnnotations error: %w", err)
	}
	return groups, token, rateLimit, nil
}

func (o *groupBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	options := []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDisplayName(fmt.Sprintf("%s Group %s", resource.DisplayName, "member")),
		ent.WithDescription(fmt.Sprintf("%s of %s Formal group", "member", resource.DisplayName)),
	}

	return []*v2.Entitlement{
		ent.NewAssignmentEntitlement(resource, "member", options...),
	}, "", nil, nil
}

func (o *groupBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	request := connect.NewRequest(&corev1.ListUserGroupLinksRequest{
		GroupId: resource.Id.Resource,
		Limit:   PageSize,
		Cursor:  pToken.Token,
	})

	response, err := o.client.GroupServiceClient.ListUserGroupLinks(ctx, request)
	if err != nil {
		return nil, "", nil, fmt.Errorf("GroupServiceClient.ListUserGroupLinks error: %w", err)
	}

	var rv []*v2.Grant
	for _, groupLink := range response.Msg.UserGroupLinks {
		getUserRequest := connect.NewRequest(&corev1.GetUserRequest{
			Id: groupLink.User.Id,
		})

		user, err := o.client.UserServiceClient.GetUser(ctx, getUserRequest)
		if err != nil {
			return nil, "", nil, fmt.Errorf("UserServiceClient.GetUser error: %w", err)
		}

		userResource, err := formalUserToResourceUser(nil, user.Msg.User)
		if err != nil {
			return nil, "", nil, fmt.Errorf("formalUserToResourceUser error: %w", err)
		}
		if userResource == nil {
			continue
		}

		rv = append(rv, grant.NewGrant(
			resource,
			"member",
			userResource,
			grant.WithExternalPrincipalID(&v2.ExternalId{Id: groupLink.Id}),
		))
	}

	token := ""
	if len(response.Msg.UserGroupLinks) > 0 {
		token = response.Msg.UserGroupLinks[len(response.Msg.UserGroupLinks)-1].Id
	}

	rateLimit, err := rateLimitAnnotations(response.Header())
	if err != nil {
		return nil, "", nil, fmt.Errorf("rateLimitAnnotations error: %w", err)
	}
	return rv, token, rateLimit, nil
}

func (o *groupBuilder) Grant(ctx context.Context, principal *v2.Resource, entitlement *v2.Entitlement) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	if principal.Id.ResourceType != userResourceType.Id {
		return nil, fmt.Errorf("only users can have group link granted")
	}

	request := connect.NewRequest(&corev1.CreateUserGroupLinkRequest{
		GroupId: entitlement.Resource.Id.Resource,
		UserId:  principal.Id.Resource,
	})

	response, err := o.client.GroupServiceClient.CreateUserGroupLink(ctx, request)
	if err != nil {
		var connectErr *connect.Error
		if errors.As(err, &connectErr) {
			if connectErr.Code() == connect.CodeAlreadyExists {
				l.Debug(
					"group link already exists, returning successfully",
					zap.String("principal", principal.Id.Resource),
					zap.String("entitlement", entitlement.Resource.Id.Resource),
				)
				return nil, nil
			}
		}

		return nil, fmt.Errorf("GroupServiceClient.CreateUserGroupLink error: %w", err)
	}

	rateLimit, err := rateLimitAnnotations(response.Header())
	if err != nil {
		return nil, fmt.Errorf("rateLimitAnnotations error: %w", err)
	}
	return rateLimit, nil
}

func (o *groupBuilder) Revoke(ctx context.Context, grant *v2.Grant) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	if grant.Principal.Id.ResourceType != userResourceType.Id {
		return nil, fmt.Errorf("only users can have group link revoked")
	}

	npt := ""
	var lastResponse *connect.Response[corev1.ListUserGroupLinksResponse]
	for {
		request := connect.NewRequest(&corev1.ListUserGroupLinksRequest{
			GroupId: grant.Entitlement.Resource.Id.Resource,
			Limit:   PageSize,
			Cursor:  npt,
		})

		response, err := o.client.GroupServiceClient.ListUserGroupLinks(ctx, request)
		if err != nil {
			return nil, fmt.Errorf("GroupServiceClient.ListUserGroupLinks error: %w", err)
		}

		lastResponse = response

		if len(response.Msg.UserGroupLinks) == 0 {
			break
		}

		for _, groupLink := range response.Msg.UserGroupLinks {
			if groupLink.User.Id == grant.Principal.Id.Resource {
				request := connect.NewRequest(&corev1.DeleteUserGroupLinkRequest{
					Id: groupLink.Id,
				})

				response, err := o.client.GroupServiceClient.DeleteUserGroupLink(ctx, request)
				if err != nil {
					return nil, fmt.Errorf("GroupServiceClient.DeleteUserGroupLink error: %w", err)
				}

				rateLimit, err := rateLimitAnnotations(response.Header())
				if err != nil {
					return nil, fmt.Errorf("rateLimitAnnotations error: %w", err)
				}
				return rateLimit, nil
			}
		}
		npt = response.Msg.UserGroupLinks[len(response.Msg.UserGroupLinks)-1].Id
	}

	rateLimit, err := rateLimitAnnotations(lastResponse.Header())
	if err != nil {
		return nil, fmt.Errorf("rateLimitAnnotations error: %w", err)
	}

	l.Debug(
		"group link not found, returning successfully",
		zap.String("principal", grant.Principal.Id.Resource),
		zap.String("entitlement", grant.Entitlement.Resource.Id.Resource),
	)
	return rateLimit, nil
}

func newGroupBuilder(client *sdk.FormalSDK) *groupBuilder {
	return &groupBuilder{
		client: client,
	}
}
