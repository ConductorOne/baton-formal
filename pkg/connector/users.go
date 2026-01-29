package connector

import (
	"context"
	"fmt"

	corev1 "buf.build/gen/go/formal/core/protocolbuffers/go/core/v1"
	"connectrpc.com/connect"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/formalco/go-sdk/sdk/v2"
)

type userBuilder struct {
	client *sdk.FormalSDK
}

func (o *userBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return userResourceType
}

func (o *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, attrs rs.SyncOpAttrs) ([]*v2.Resource, *rs.SyncOpResults, error) {
	pToken := attrs.PageToken
	request := connect.NewRequest(&corev1.ListUsersRequest{
		Limit:  100,
		Cursor: pToken.Token,
	})

	response, err := o.client.UserServiceClient.ListUsers(ctx, request)
	if err != nil {
		return nil, nil, fmt.Errorf("UserServiceClient.ListUsers error: %w", err)
	}

	var users []*v2.Resource
	for _, user := range response.Msg.Users {
		userResource, err := formalUserToResourceUser(parentResourceID, user)
		if err != nil {
			return nil, nil, fmt.Errorf("formalUserToResourceUser error: %w", err)
		}
		if userResource == nil {
			continue
		}
		users = append(users, userResource)
	}

	token := ""
	if len(response.Msg.Users) > 0 {
		token = response.Msg.Users[len(response.Msg.Users)-1].Id
	}

	rateLimit, err := rateLimitAnnotations(response.Header())
	if err != nil {
		return nil, nil, fmt.Errorf("rateLimitAnnotations error: %w", err)
	}
	return users, &rs.SyncOpResults{NextPageToken: token, Annotations: rateLimit}, nil
}

// Entitlements always returns an empty slice for users.
func (o *userBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ rs.SyncOpAttrs) ([]*v2.Entitlement, *rs.SyncOpResults, error) {
	return nil, nil, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (o *userBuilder) Grants(ctx context.Context, resource *v2.Resource, attrs rs.SyncOpAttrs) ([]*v2.Grant, *rs.SyncOpResults, error) {
	return nil, nil, nil
}

func newUserBuilder(client *sdk.FormalSDK) *userBuilder {
	return &userBuilder{
		client: client,
	}
}
