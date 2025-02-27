// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: core/v1/user.proto

package corev1connect

import (
	v1 "buf.build/gen/go/formal/core/protocolbuffers/go/core/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "core.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceCreateUserProcedure is the fully-qualified name of the UserService's CreateUser RPC.
	UserServiceCreateUserProcedure = "/core.v1.UserService/CreateUser"
	// UserServiceDeleteUserProcedure is the fully-qualified name of the UserService's DeleteUser RPC.
	UserServiceDeleteUserProcedure = "/core.v1.UserService/DeleteUser"
	// UserServiceUpdateUserProcedure is the fully-qualified name of the UserService's UpdateUser RPC.
	UserServiceUpdateUserProcedure = "/core.v1.UserService/UpdateUser"
	// UserServiceListUsersProcedure is the fully-qualified name of the UserService's ListUsers RPC.
	UserServiceListUsersProcedure = "/core.v1.UserService/ListUsers"
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/core.v1.UserService/GetUser"
	// UserServiceGetMachineUserCredentialsProcedure is the fully-qualified name of the UserService's
	// GetMachineUserCredentials RPC.
	UserServiceGetMachineUserCredentialsProcedure = "/core.v1.UserService/GetMachineUserCredentials"
	// UserServiceCreateMachineUserPasswordProcedure is the fully-qualified name of the UserService's
	// CreateMachineUserPassword RPC.
	UserServiceCreateMachineUserPasswordProcedure = "/core.v1.UserService/CreateMachineUserPassword"
	// UserServiceCreateHumanUserPasswordProcedure is the fully-qualified name of the UserService's
	// CreateHumanUserPassword RPC.
	UserServiceCreateHumanUserPasswordProcedure = "/core.v1.UserService/CreateHumanUserPassword"
	// UserServiceListUserExternalIdsProcedure is the fully-qualified name of the UserService's
	// ListUserExternalIds RPC.
	UserServiceListUserExternalIdsProcedure = "/core.v1.UserService/ListUserExternalIds"
	// UserServiceCreateUserExternalIdProcedure is the fully-qualified name of the UserService's
	// CreateUserExternalId RPC.
	UserServiceCreateUserExternalIdProcedure = "/core.v1.UserService/CreateUserExternalId"
	// UserServiceDeleteUserExternalIdProcedure is the fully-qualified name of the UserService's
	// DeleteUserExternalId RPC.
	UserServiceDeleteUserExternalIdProcedure = "/core.v1.UserService/DeleteUserExternalId"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor                         = v1.File_core_v1_user_proto.Services().ByName("UserService")
	userServiceCreateUserMethodDescriptor                = userServiceServiceDescriptor.Methods().ByName("CreateUser")
	userServiceDeleteUserMethodDescriptor                = userServiceServiceDescriptor.Methods().ByName("DeleteUser")
	userServiceUpdateUserMethodDescriptor                = userServiceServiceDescriptor.Methods().ByName("UpdateUser")
	userServiceListUsersMethodDescriptor                 = userServiceServiceDescriptor.Methods().ByName("ListUsers")
	userServiceGetUserMethodDescriptor                   = userServiceServiceDescriptor.Methods().ByName("GetUser")
	userServiceGetMachineUserCredentialsMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("GetMachineUserCredentials")
	userServiceCreateMachineUserPasswordMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("CreateMachineUserPassword")
	userServiceCreateHumanUserPasswordMethodDescriptor   = userServiceServiceDescriptor.Methods().ByName("CreateHumanUserPassword")
	userServiceListUserExternalIdsMethodDescriptor       = userServiceServiceDescriptor.Methods().ByName("ListUserExternalIds")
	userServiceCreateUserExternalIdMethodDescriptor      = userServiceServiceDescriptor.Methods().ByName("CreateUserExternalId")
	userServiceDeleteUserExternalIdMethodDescriptor      = userServiceServiceDescriptor.Methods().ByName("DeleteUserExternalId")
)

// UserServiceClient is a client for the core.v1.UserService service.
type UserServiceClient interface {
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error)
	UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error)
	ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error)
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	GetMachineUserCredentials(context.Context, *connect.Request[v1.GetMachineUserCredentialsRequest]) (*connect.Response[v1.GetMachineUserCredentialsResponse], error)
	CreateMachineUserPassword(context.Context, *connect.Request[v1.CreateMachineUserPasswordRequest]) (*connect.Response[v1.CreateMachineUserPasswordResponse], error)
	CreateHumanUserPassword(context.Context, *connect.Request[v1.CreateHumanUserPasswordRequest]) (*connect.Response[v1.CreateHumanUserPasswordResponse], error)
	ListUserExternalIds(context.Context, *connect.Request[v1.ListUserExternalIdsRequest]) (*connect.Response[v1.ListUserExternalIdsResponse], error)
	CreateUserExternalId(context.Context, *connect.Request[v1.CreateUserExternalIdRequest]) (*connect.Response[v1.CreateUserExternalIdResponse], error)
	DeleteUserExternalId(context.Context, *connect.Request[v1.DeleteUserExternalIdRequest]) (*connect.Response[v1.DeleteUserExternalIdResponse], error)
}

// NewUserServiceClient constructs a client for the core.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		createUser: connect.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+UserServiceCreateUserProcedure,
			connect.WithSchema(userServiceCreateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUser: connect.NewClient[v1.DeleteUserRequest, v1.DeleteUserResponse](
			httpClient,
			baseURL+UserServiceDeleteUserProcedure,
			connect.WithSchema(userServiceDeleteUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateUser: connect.NewClient[v1.UpdateUserRequest, v1.UpdateUserResponse](
			httpClient,
			baseURL+UserServiceUpdateUserProcedure,
			connect.WithSchema(userServiceUpdateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listUsers: connect.NewClient[v1.ListUsersRequest, v1.ListUsersResponse](
			httpClient,
			baseURL+UserServiceListUsersProcedure,
			connect.WithSchema(userServiceListUsersMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getUser: connect.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			connect.WithSchema(userServiceGetUserMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getMachineUserCredentials: connect.NewClient[v1.GetMachineUserCredentialsRequest, v1.GetMachineUserCredentialsResponse](
			httpClient,
			baseURL+UserServiceGetMachineUserCredentialsProcedure,
			connect.WithSchema(userServiceGetMachineUserCredentialsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createMachineUserPassword: connect.NewClient[v1.CreateMachineUserPasswordRequest, v1.CreateMachineUserPasswordResponse](
			httpClient,
			baseURL+UserServiceCreateMachineUserPasswordProcedure,
			connect.WithSchema(userServiceCreateMachineUserPasswordMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createHumanUserPassword: connect.NewClient[v1.CreateHumanUserPasswordRequest, v1.CreateHumanUserPasswordResponse](
			httpClient,
			baseURL+UserServiceCreateHumanUserPasswordProcedure,
			connect.WithSchema(userServiceCreateHumanUserPasswordMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listUserExternalIds: connect.NewClient[v1.ListUserExternalIdsRequest, v1.ListUserExternalIdsResponse](
			httpClient,
			baseURL+UserServiceListUserExternalIdsProcedure,
			connect.WithSchema(userServiceListUserExternalIdsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createUserExternalId: connect.NewClient[v1.CreateUserExternalIdRequest, v1.CreateUserExternalIdResponse](
			httpClient,
			baseURL+UserServiceCreateUserExternalIdProcedure,
			connect.WithSchema(userServiceCreateUserExternalIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUserExternalId: connect.NewClient[v1.DeleteUserExternalIdRequest, v1.DeleteUserExternalIdResponse](
			httpClient,
			baseURL+UserServiceDeleteUserExternalIdProcedure,
			connect.WithSchema(userServiceDeleteUserExternalIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	createUser                *connect.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	deleteUser                *connect.Client[v1.DeleteUserRequest, v1.DeleteUserResponse]
	updateUser                *connect.Client[v1.UpdateUserRequest, v1.UpdateUserResponse]
	listUsers                 *connect.Client[v1.ListUsersRequest, v1.ListUsersResponse]
	getUser                   *connect.Client[v1.GetUserRequest, v1.GetUserResponse]
	getMachineUserCredentials *connect.Client[v1.GetMachineUserCredentialsRequest, v1.GetMachineUserCredentialsResponse]
	createMachineUserPassword *connect.Client[v1.CreateMachineUserPasswordRequest, v1.CreateMachineUserPasswordResponse]
	createHumanUserPassword   *connect.Client[v1.CreateHumanUserPasswordRequest, v1.CreateHumanUserPasswordResponse]
	listUserExternalIds       *connect.Client[v1.ListUserExternalIdsRequest, v1.ListUserExternalIdsResponse]
	createUserExternalId      *connect.Client[v1.CreateUserExternalIdRequest, v1.CreateUserExternalIdResponse]
	deleteUserExternalId      *connect.Client[v1.DeleteUserExternalIdRequest, v1.DeleteUserExternalIdResponse]
}

// CreateUser calls core.v1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// DeleteUser calls core.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// UpdateUser calls core.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// ListUsers calls core.v1.UserService.ListUsers.
func (c *userServiceClient) ListUsers(ctx context.Context, req *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// GetUser calls core.v1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// GetMachineUserCredentials calls core.v1.UserService.GetMachineUserCredentials.
func (c *userServiceClient) GetMachineUserCredentials(ctx context.Context, req *connect.Request[v1.GetMachineUserCredentialsRequest]) (*connect.Response[v1.GetMachineUserCredentialsResponse], error) {
	return c.getMachineUserCredentials.CallUnary(ctx, req)
}

// CreateMachineUserPassword calls core.v1.UserService.CreateMachineUserPassword.
func (c *userServiceClient) CreateMachineUserPassword(ctx context.Context, req *connect.Request[v1.CreateMachineUserPasswordRequest]) (*connect.Response[v1.CreateMachineUserPasswordResponse], error) {
	return c.createMachineUserPassword.CallUnary(ctx, req)
}

// CreateHumanUserPassword calls core.v1.UserService.CreateHumanUserPassword.
func (c *userServiceClient) CreateHumanUserPassword(ctx context.Context, req *connect.Request[v1.CreateHumanUserPasswordRequest]) (*connect.Response[v1.CreateHumanUserPasswordResponse], error) {
	return c.createHumanUserPassword.CallUnary(ctx, req)
}

// ListUserExternalIds calls core.v1.UserService.ListUserExternalIds.
func (c *userServiceClient) ListUserExternalIds(ctx context.Context, req *connect.Request[v1.ListUserExternalIdsRequest]) (*connect.Response[v1.ListUserExternalIdsResponse], error) {
	return c.listUserExternalIds.CallUnary(ctx, req)
}

// CreateUserExternalId calls core.v1.UserService.CreateUserExternalId.
func (c *userServiceClient) CreateUserExternalId(ctx context.Context, req *connect.Request[v1.CreateUserExternalIdRequest]) (*connect.Response[v1.CreateUserExternalIdResponse], error) {
	return c.createUserExternalId.CallUnary(ctx, req)
}

// DeleteUserExternalId calls core.v1.UserService.DeleteUserExternalId.
func (c *userServiceClient) DeleteUserExternalId(ctx context.Context, req *connect.Request[v1.DeleteUserExternalIdRequest]) (*connect.Response[v1.DeleteUserExternalIdResponse], error) {
	return c.deleteUserExternalId.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the core.v1.UserService service.
type UserServiceHandler interface {
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error)
	UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error)
	ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error)
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	GetMachineUserCredentials(context.Context, *connect.Request[v1.GetMachineUserCredentialsRequest]) (*connect.Response[v1.GetMachineUserCredentialsResponse], error)
	CreateMachineUserPassword(context.Context, *connect.Request[v1.CreateMachineUserPasswordRequest]) (*connect.Response[v1.CreateMachineUserPasswordResponse], error)
	CreateHumanUserPassword(context.Context, *connect.Request[v1.CreateHumanUserPasswordRequest]) (*connect.Response[v1.CreateHumanUserPasswordResponse], error)
	ListUserExternalIds(context.Context, *connect.Request[v1.ListUserExternalIdsRequest]) (*connect.Response[v1.ListUserExternalIdsResponse], error)
	CreateUserExternalId(context.Context, *connect.Request[v1.CreateUserExternalIdRequest]) (*connect.Response[v1.CreateUserExternalIdResponse], error)
	DeleteUserExternalId(context.Context, *connect.Request[v1.DeleteUserExternalIdRequest]) (*connect.Response[v1.DeleteUserExternalIdResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceCreateUserHandler := connect.NewUnaryHandler(
		UserServiceCreateUserProcedure,
		svc.CreateUser,
		connect.WithSchema(userServiceCreateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteUserHandler := connect.NewUnaryHandler(
		UserServiceDeleteUserProcedure,
		svc.DeleteUser,
		connect.WithSchema(userServiceDeleteUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceUpdateUserHandler := connect.NewUnaryHandler(
		UserServiceUpdateUserProcedure,
		svc.UpdateUser,
		connect.WithSchema(userServiceUpdateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceListUsersHandler := connect.NewUnaryHandler(
		UserServiceListUsersProcedure,
		svc.ListUsers,
		connect.WithSchema(userServiceListUsersMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetUserHandler := connect.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(userServiceGetUserMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetMachineUserCredentialsHandler := connect.NewUnaryHandler(
		UserServiceGetMachineUserCredentialsProcedure,
		svc.GetMachineUserCredentials,
		connect.WithSchema(userServiceGetMachineUserCredentialsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateMachineUserPasswordHandler := connect.NewUnaryHandler(
		UserServiceCreateMachineUserPasswordProcedure,
		svc.CreateMachineUserPassword,
		connect.WithSchema(userServiceCreateMachineUserPasswordMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateHumanUserPasswordHandler := connect.NewUnaryHandler(
		UserServiceCreateHumanUserPasswordProcedure,
		svc.CreateHumanUserPassword,
		connect.WithSchema(userServiceCreateHumanUserPasswordMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceListUserExternalIdsHandler := connect.NewUnaryHandler(
		UserServiceListUserExternalIdsProcedure,
		svc.ListUserExternalIds,
		connect.WithSchema(userServiceListUserExternalIdsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateUserExternalIdHandler := connect.NewUnaryHandler(
		UserServiceCreateUserExternalIdProcedure,
		svc.CreateUserExternalId,
		connect.WithSchema(userServiceCreateUserExternalIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteUserExternalIdHandler := connect.NewUnaryHandler(
		UserServiceDeleteUserExternalIdProcedure,
		svc.DeleteUserExternalId,
		connect.WithSchema(userServiceDeleteUserExternalIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/core.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceCreateUserProcedure:
			userServiceCreateUserHandler.ServeHTTP(w, r)
		case UserServiceDeleteUserProcedure:
			userServiceDeleteUserHandler.ServeHTTP(w, r)
		case UserServiceUpdateUserProcedure:
			userServiceUpdateUserHandler.ServeHTTP(w, r)
		case UserServiceListUsersProcedure:
			userServiceListUsersHandler.ServeHTTP(w, r)
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		case UserServiceGetMachineUserCredentialsProcedure:
			userServiceGetMachineUserCredentialsHandler.ServeHTTP(w, r)
		case UserServiceCreateMachineUserPasswordProcedure:
			userServiceCreateMachineUserPasswordHandler.ServeHTTP(w, r)
		case UserServiceCreateHumanUserPasswordProcedure:
			userServiceCreateHumanUserPasswordHandler.ServeHTTP(w, r)
		case UserServiceListUserExternalIdsProcedure:
			userServiceListUserExternalIdsHandler.ServeHTTP(w, r)
		case UserServiceCreateUserExternalIdProcedure:
			userServiceCreateUserExternalIdHandler.ServeHTTP(w, r)
		case UserServiceDeleteUserExternalIdProcedure:
			userServiceDeleteUserExternalIdHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.DeleteUser is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.ListUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetMachineUserCredentials(context.Context, *connect.Request[v1.GetMachineUserCredentialsRequest]) (*connect.Response[v1.GetMachineUserCredentialsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.GetMachineUserCredentials is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateMachineUserPassword(context.Context, *connect.Request[v1.CreateMachineUserPasswordRequest]) (*connect.Response[v1.CreateMachineUserPasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.CreateMachineUserPassword is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateHumanUserPassword(context.Context, *connect.Request[v1.CreateHumanUserPasswordRequest]) (*connect.Response[v1.CreateHumanUserPasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.CreateHumanUserPassword is not implemented"))
}

func (UnimplementedUserServiceHandler) ListUserExternalIds(context.Context, *connect.Request[v1.ListUserExternalIdsRequest]) (*connect.Response[v1.ListUserExternalIdsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.ListUserExternalIds is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateUserExternalId(context.Context, *connect.Request[v1.CreateUserExternalIdRequest]) (*connect.Response[v1.CreateUserExternalIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.CreateUserExternalId is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUserExternalId(context.Context, *connect.Request[v1.DeleteUserExternalIdRequest]) (*connect.Response[v1.DeleteUserExternalIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.UserService.DeleteUserExternalId is not implemented"))
}
