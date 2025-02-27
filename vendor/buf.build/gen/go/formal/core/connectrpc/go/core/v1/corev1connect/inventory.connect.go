// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: core/v1/inventory.proto

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
	// InventoryServiceName is the fully-qualified name of the InventoryService service.
	InventoryServiceName = "core.v1.InventoryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// InventoryServiceUpdateColumnProcedure is the fully-qualified name of the InventoryService's
	// UpdateColumn RPC.
	InventoryServiceUpdateColumnProcedure = "/core.v1.InventoryService/UpdateColumn"
	// InventoryServiceCreateInventoryObjectProcedure is the fully-qualified name of the
	// InventoryService's CreateInventoryObject RPC.
	InventoryServiceCreateInventoryObjectProcedure = "/core.v1.InventoryService/CreateInventoryObject"
	// InventoryServiceUpdateInventoryObjectTagsProcedure is the fully-qualified name of the
	// InventoryService's UpdateInventoryObjectTags RPC.
	InventoryServiceUpdateInventoryObjectTagsProcedure = "/core.v1.InventoryService/UpdateInventoryObjectTags"
	// InventoryServiceDeleteInventoryObjectProcedure is the fully-qualified name of the
	// InventoryService's DeleteInventoryObject RPC.
	InventoryServiceDeleteInventoryObjectProcedure = "/core.v1.InventoryService/DeleteInventoryObject"
	// InventoryServiceGetInventoryObjectProcedure is the fully-qualified name of the InventoryService's
	// GetInventoryObject RPC.
	InventoryServiceGetInventoryObjectProcedure = "/core.v1.InventoryService/GetInventoryObject"
	// InventoryServiceListInventoryObjectsProcedure is the fully-qualified name of the
	// InventoryService's ListInventoryObjects RPC.
	InventoryServiceListInventoryObjectsProcedure = "/core.v1.InventoryService/ListInventoryObjects"
	// InventoryServiceCreateInventoryTagProcedure is the fully-qualified name of the InventoryService's
	// CreateInventoryTag RPC.
	InventoryServiceCreateInventoryTagProcedure = "/core.v1.InventoryService/CreateInventoryTag"
	// InventoryServiceListInventoryTagsProcedure is the fully-qualified name of the InventoryService's
	// ListInventoryTags RPC.
	InventoryServiceListInventoryTagsProcedure = "/core.v1.InventoryService/ListInventoryTags"
	// InventoryServiceDeleteInventoryTagProcedure is the fully-qualified name of the InventoryService's
	// DeleteInventoryTag RPC.
	InventoryServiceDeleteInventoryTagProcedure = "/core.v1.InventoryService/DeleteInventoryTag"
	// InventoryServiceCreateDataDomainProcedure is the fully-qualified name of the InventoryService's
	// CreateDataDomain RPC.
	InventoryServiceCreateDataDomainProcedure = "/core.v1.InventoryService/CreateDataDomain"
	// InventoryServiceListDataDomainsProcedure is the fully-qualified name of the InventoryService's
	// ListDataDomains RPC.
	InventoryServiceListDataDomainsProcedure = "/core.v1.InventoryService/ListDataDomains"
	// InventoryServiceGetDataDomainProcedure is the fully-qualified name of the InventoryService's
	// GetDataDomain RPC.
	InventoryServiceGetDataDomainProcedure = "/core.v1.InventoryService/GetDataDomain"
	// InventoryServiceUpdateDataDomainProcedure is the fully-qualified name of the InventoryService's
	// UpdateDataDomain RPC.
	InventoryServiceUpdateDataDomainProcedure = "/core.v1.InventoryService/UpdateDataDomain"
	// InventoryServiceDeleteDataDomainProcedure is the fully-qualified name of the InventoryService's
	// DeleteDataDomain RPC.
	InventoryServiceDeleteDataDomainProcedure = "/core.v1.InventoryService/DeleteDataDomain"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	inventoryServiceServiceDescriptor                         = v1.File_core_v1_inventory_proto.Services().ByName("InventoryService")
	inventoryServiceUpdateColumnMethodDescriptor              = inventoryServiceServiceDescriptor.Methods().ByName("UpdateColumn")
	inventoryServiceCreateInventoryObjectMethodDescriptor     = inventoryServiceServiceDescriptor.Methods().ByName("CreateInventoryObject")
	inventoryServiceUpdateInventoryObjectTagsMethodDescriptor = inventoryServiceServiceDescriptor.Methods().ByName("UpdateInventoryObjectTags")
	inventoryServiceDeleteInventoryObjectMethodDescriptor     = inventoryServiceServiceDescriptor.Methods().ByName("DeleteInventoryObject")
	inventoryServiceGetInventoryObjectMethodDescriptor        = inventoryServiceServiceDescriptor.Methods().ByName("GetInventoryObject")
	inventoryServiceListInventoryObjectsMethodDescriptor      = inventoryServiceServiceDescriptor.Methods().ByName("ListInventoryObjects")
	inventoryServiceCreateInventoryTagMethodDescriptor        = inventoryServiceServiceDescriptor.Methods().ByName("CreateInventoryTag")
	inventoryServiceListInventoryTagsMethodDescriptor         = inventoryServiceServiceDescriptor.Methods().ByName("ListInventoryTags")
	inventoryServiceDeleteInventoryTagMethodDescriptor        = inventoryServiceServiceDescriptor.Methods().ByName("DeleteInventoryTag")
	inventoryServiceCreateDataDomainMethodDescriptor          = inventoryServiceServiceDescriptor.Methods().ByName("CreateDataDomain")
	inventoryServiceListDataDomainsMethodDescriptor           = inventoryServiceServiceDescriptor.Methods().ByName("ListDataDomains")
	inventoryServiceGetDataDomainMethodDescriptor             = inventoryServiceServiceDescriptor.Methods().ByName("GetDataDomain")
	inventoryServiceUpdateDataDomainMethodDescriptor          = inventoryServiceServiceDescriptor.Methods().ByName("UpdateDataDomain")
	inventoryServiceDeleteDataDomainMethodDescriptor          = inventoryServiceServiceDescriptor.Methods().ByName("DeleteDataDomain")
)

// InventoryServiceClient is a client for the core.v1.InventoryService service.
type InventoryServiceClient interface {
	UpdateColumn(context.Context, *connect.Request[v1.UpdateColumnRequest]) (*connect.Response[v1.UpdateColumnResponse], error)
	CreateInventoryObject(context.Context, *connect.Request[v1.CreateInventoryObjectRequest]) (*connect.Response[v1.CreateInventoryObjectResponse], error)
	UpdateInventoryObjectTags(context.Context, *connect.Request[v1.UpdateInventoryObjectTagsRequest]) (*connect.Response[v1.UpdateInventoryObjectTagsResponse], error)
	DeleteInventoryObject(context.Context, *connect.Request[v1.DeleteInventoryObjectRequest]) (*connect.Response[v1.DeleteInventoryObjectResponse], error)
	GetInventoryObject(context.Context, *connect.Request[v1.GetInventoryObjectRequest]) (*connect.Response[v1.GetInventoryObjectResponse], error)
	ListInventoryObjects(context.Context, *connect.Request[v1.ListInventoryObjectsRequest]) (*connect.Response[v1.ListInventoryObjectsResponse], error)
	CreateInventoryTag(context.Context, *connect.Request[v1.CreateInventoryTagRequest]) (*connect.Response[v1.CreateInventoryTagResponse], error)
	ListInventoryTags(context.Context, *connect.Request[v1.ListInventoryTagsRequest]) (*connect.Response[v1.ListInventoryTagsResponse], error)
	DeleteInventoryTag(context.Context, *connect.Request[v1.DeleteInventoryTagRequest]) (*connect.Response[v1.DeleteInventoryTagResponse], error)
	CreateDataDomain(context.Context, *connect.Request[v1.CreateDataDomainRequest]) (*connect.Response[v1.CreateDataDomainResponse], error)
	ListDataDomains(context.Context, *connect.Request[v1.ListDataDomainsRequest]) (*connect.Response[v1.ListDataDomainsResponse], error)
	GetDataDomain(context.Context, *connect.Request[v1.GetDataDomainRequest]) (*connect.Response[v1.GetDataDomainResponse], error)
	UpdateDataDomain(context.Context, *connect.Request[v1.UpdateDataDomainRequest]) (*connect.Response[v1.UpdateDataDomainResponse], error)
	DeleteDataDomain(context.Context, *connect.Request[v1.DeleteDataDomainRequest]) (*connect.Response[v1.DeleteDataDomainResponse], error)
}

// NewInventoryServiceClient constructs a client for the core.v1.InventoryService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewInventoryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) InventoryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &inventoryServiceClient{
		updateColumn: connect.NewClient[v1.UpdateColumnRequest, v1.UpdateColumnResponse](
			httpClient,
			baseURL+InventoryServiceUpdateColumnProcedure,
			connect.WithSchema(inventoryServiceUpdateColumnMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createInventoryObject: connect.NewClient[v1.CreateInventoryObjectRequest, v1.CreateInventoryObjectResponse](
			httpClient,
			baseURL+InventoryServiceCreateInventoryObjectProcedure,
			connect.WithSchema(inventoryServiceCreateInventoryObjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateInventoryObjectTags: connect.NewClient[v1.UpdateInventoryObjectTagsRequest, v1.UpdateInventoryObjectTagsResponse](
			httpClient,
			baseURL+InventoryServiceUpdateInventoryObjectTagsProcedure,
			connect.WithSchema(inventoryServiceUpdateInventoryObjectTagsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteInventoryObject: connect.NewClient[v1.DeleteInventoryObjectRequest, v1.DeleteInventoryObjectResponse](
			httpClient,
			baseURL+InventoryServiceDeleteInventoryObjectProcedure,
			connect.WithSchema(inventoryServiceDeleteInventoryObjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getInventoryObject: connect.NewClient[v1.GetInventoryObjectRequest, v1.GetInventoryObjectResponse](
			httpClient,
			baseURL+InventoryServiceGetInventoryObjectProcedure,
			connect.WithSchema(inventoryServiceGetInventoryObjectMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listInventoryObjects: connect.NewClient[v1.ListInventoryObjectsRequest, v1.ListInventoryObjectsResponse](
			httpClient,
			baseURL+InventoryServiceListInventoryObjectsProcedure,
			connect.WithSchema(inventoryServiceListInventoryObjectsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createInventoryTag: connect.NewClient[v1.CreateInventoryTagRequest, v1.CreateInventoryTagResponse](
			httpClient,
			baseURL+InventoryServiceCreateInventoryTagProcedure,
			connect.WithSchema(inventoryServiceCreateInventoryTagMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listInventoryTags: connect.NewClient[v1.ListInventoryTagsRequest, v1.ListInventoryTagsResponse](
			httpClient,
			baseURL+InventoryServiceListInventoryTagsProcedure,
			connect.WithSchema(inventoryServiceListInventoryTagsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		deleteInventoryTag: connect.NewClient[v1.DeleteInventoryTagRequest, v1.DeleteInventoryTagResponse](
			httpClient,
			baseURL+InventoryServiceDeleteInventoryTagProcedure,
			connect.WithSchema(inventoryServiceDeleteInventoryTagMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createDataDomain: connect.NewClient[v1.CreateDataDomainRequest, v1.CreateDataDomainResponse](
			httpClient,
			baseURL+InventoryServiceCreateDataDomainProcedure,
			connect.WithSchema(inventoryServiceCreateDataDomainMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listDataDomains: connect.NewClient[v1.ListDataDomainsRequest, v1.ListDataDomainsResponse](
			httpClient,
			baseURL+InventoryServiceListDataDomainsProcedure,
			connect.WithSchema(inventoryServiceListDataDomainsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getDataDomain: connect.NewClient[v1.GetDataDomainRequest, v1.GetDataDomainResponse](
			httpClient,
			baseURL+InventoryServiceGetDataDomainProcedure,
			connect.WithSchema(inventoryServiceGetDataDomainMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		updateDataDomain: connect.NewClient[v1.UpdateDataDomainRequest, v1.UpdateDataDomainResponse](
			httpClient,
			baseURL+InventoryServiceUpdateDataDomainProcedure,
			connect.WithSchema(inventoryServiceUpdateDataDomainMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteDataDomain: connect.NewClient[v1.DeleteDataDomainRequest, v1.DeleteDataDomainResponse](
			httpClient,
			baseURL+InventoryServiceDeleteDataDomainProcedure,
			connect.WithSchema(inventoryServiceDeleteDataDomainMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// inventoryServiceClient implements InventoryServiceClient.
type inventoryServiceClient struct {
	updateColumn              *connect.Client[v1.UpdateColumnRequest, v1.UpdateColumnResponse]
	createInventoryObject     *connect.Client[v1.CreateInventoryObjectRequest, v1.CreateInventoryObjectResponse]
	updateInventoryObjectTags *connect.Client[v1.UpdateInventoryObjectTagsRequest, v1.UpdateInventoryObjectTagsResponse]
	deleteInventoryObject     *connect.Client[v1.DeleteInventoryObjectRequest, v1.DeleteInventoryObjectResponse]
	getInventoryObject        *connect.Client[v1.GetInventoryObjectRequest, v1.GetInventoryObjectResponse]
	listInventoryObjects      *connect.Client[v1.ListInventoryObjectsRequest, v1.ListInventoryObjectsResponse]
	createInventoryTag        *connect.Client[v1.CreateInventoryTagRequest, v1.CreateInventoryTagResponse]
	listInventoryTags         *connect.Client[v1.ListInventoryTagsRequest, v1.ListInventoryTagsResponse]
	deleteInventoryTag        *connect.Client[v1.DeleteInventoryTagRequest, v1.DeleteInventoryTagResponse]
	createDataDomain          *connect.Client[v1.CreateDataDomainRequest, v1.CreateDataDomainResponse]
	listDataDomains           *connect.Client[v1.ListDataDomainsRequest, v1.ListDataDomainsResponse]
	getDataDomain             *connect.Client[v1.GetDataDomainRequest, v1.GetDataDomainResponse]
	updateDataDomain          *connect.Client[v1.UpdateDataDomainRequest, v1.UpdateDataDomainResponse]
	deleteDataDomain          *connect.Client[v1.DeleteDataDomainRequest, v1.DeleteDataDomainResponse]
}

// UpdateColumn calls core.v1.InventoryService.UpdateColumn.
func (c *inventoryServiceClient) UpdateColumn(ctx context.Context, req *connect.Request[v1.UpdateColumnRequest]) (*connect.Response[v1.UpdateColumnResponse], error) {
	return c.updateColumn.CallUnary(ctx, req)
}

// CreateInventoryObject calls core.v1.InventoryService.CreateInventoryObject.
func (c *inventoryServiceClient) CreateInventoryObject(ctx context.Context, req *connect.Request[v1.CreateInventoryObjectRequest]) (*connect.Response[v1.CreateInventoryObjectResponse], error) {
	return c.createInventoryObject.CallUnary(ctx, req)
}

// UpdateInventoryObjectTags calls core.v1.InventoryService.UpdateInventoryObjectTags.
func (c *inventoryServiceClient) UpdateInventoryObjectTags(ctx context.Context, req *connect.Request[v1.UpdateInventoryObjectTagsRequest]) (*connect.Response[v1.UpdateInventoryObjectTagsResponse], error) {
	return c.updateInventoryObjectTags.CallUnary(ctx, req)
}

// DeleteInventoryObject calls core.v1.InventoryService.DeleteInventoryObject.
func (c *inventoryServiceClient) DeleteInventoryObject(ctx context.Context, req *connect.Request[v1.DeleteInventoryObjectRequest]) (*connect.Response[v1.DeleteInventoryObjectResponse], error) {
	return c.deleteInventoryObject.CallUnary(ctx, req)
}

// GetInventoryObject calls core.v1.InventoryService.GetInventoryObject.
func (c *inventoryServiceClient) GetInventoryObject(ctx context.Context, req *connect.Request[v1.GetInventoryObjectRequest]) (*connect.Response[v1.GetInventoryObjectResponse], error) {
	return c.getInventoryObject.CallUnary(ctx, req)
}

// ListInventoryObjects calls core.v1.InventoryService.ListInventoryObjects.
func (c *inventoryServiceClient) ListInventoryObjects(ctx context.Context, req *connect.Request[v1.ListInventoryObjectsRequest]) (*connect.Response[v1.ListInventoryObjectsResponse], error) {
	return c.listInventoryObjects.CallUnary(ctx, req)
}

// CreateInventoryTag calls core.v1.InventoryService.CreateInventoryTag.
func (c *inventoryServiceClient) CreateInventoryTag(ctx context.Context, req *connect.Request[v1.CreateInventoryTagRequest]) (*connect.Response[v1.CreateInventoryTagResponse], error) {
	return c.createInventoryTag.CallUnary(ctx, req)
}

// ListInventoryTags calls core.v1.InventoryService.ListInventoryTags.
func (c *inventoryServiceClient) ListInventoryTags(ctx context.Context, req *connect.Request[v1.ListInventoryTagsRequest]) (*connect.Response[v1.ListInventoryTagsResponse], error) {
	return c.listInventoryTags.CallUnary(ctx, req)
}

// DeleteInventoryTag calls core.v1.InventoryService.DeleteInventoryTag.
func (c *inventoryServiceClient) DeleteInventoryTag(ctx context.Context, req *connect.Request[v1.DeleteInventoryTagRequest]) (*connect.Response[v1.DeleteInventoryTagResponse], error) {
	return c.deleteInventoryTag.CallUnary(ctx, req)
}

// CreateDataDomain calls core.v1.InventoryService.CreateDataDomain.
func (c *inventoryServiceClient) CreateDataDomain(ctx context.Context, req *connect.Request[v1.CreateDataDomainRequest]) (*connect.Response[v1.CreateDataDomainResponse], error) {
	return c.createDataDomain.CallUnary(ctx, req)
}

// ListDataDomains calls core.v1.InventoryService.ListDataDomains.
func (c *inventoryServiceClient) ListDataDomains(ctx context.Context, req *connect.Request[v1.ListDataDomainsRequest]) (*connect.Response[v1.ListDataDomainsResponse], error) {
	return c.listDataDomains.CallUnary(ctx, req)
}

// GetDataDomain calls core.v1.InventoryService.GetDataDomain.
func (c *inventoryServiceClient) GetDataDomain(ctx context.Context, req *connect.Request[v1.GetDataDomainRequest]) (*connect.Response[v1.GetDataDomainResponse], error) {
	return c.getDataDomain.CallUnary(ctx, req)
}

// UpdateDataDomain calls core.v1.InventoryService.UpdateDataDomain.
func (c *inventoryServiceClient) UpdateDataDomain(ctx context.Context, req *connect.Request[v1.UpdateDataDomainRequest]) (*connect.Response[v1.UpdateDataDomainResponse], error) {
	return c.updateDataDomain.CallUnary(ctx, req)
}

// DeleteDataDomain calls core.v1.InventoryService.DeleteDataDomain.
func (c *inventoryServiceClient) DeleteDataDomain(ctx context.Context, req *connect.Request[v1.DeleteDataDomainRequest]) (*connect.Response[v1.DeleteDataDomainResponse], error) {
	return c.deleteDataDomain.CallUnary(ctx, req)
}

// InventoryServiceHandler is an implementation of the core.v1.InventoryService service.
type InventoryServiceHandler interface {
	UpdateColumn(context.Context, *connect.Request[v1.UpdateColumnRequest]) (*connect.Response[v1.UpdateColumnResponse], error)
	CreateInventoryObject(context.Context, *connect.Request[v1.CreateInventoryObjectRequest]) (*connect.Response[v1.CreateInventoryObjectResponse], error)
	UpdateInventoryObjectTags(context.Context, *connect.Request[v1.UpdateInventoryObjectTagsRequest]) (*connect.Response[v1.UpdateInventoryObjectTagsResponse], error)
	DeleteInventoryObject(context.Context, *connect.Request[v1.DeleteInventoryObjectRequest]) (*connect.Response[v1.DeleteInventoryObjectResponse], error)
	GetInventoryObject(context.Context, *connect.Request[v1.GetInventoryObjectRequest]) (*connect.Response[v1.GetInventoryObjectResponse], error)
	ListInventoryObjects(context.Context, *connect.Request[v1.ListInventoryObjectsRequest]) (*connect.Response[v1.ListInventoryObjectsResponse], error)
	CreateInventoryTag(context.Context, *connect.Request[v1.CreateInventoryTagRequest]) (*connect.Response[v1.CreateInventoryTagResponse], error)
	ListInventoryTags(context.Context, *connect.Request[v1.ListInventoryTagsRequest]) (*connect.Response[v1.ListInventoryTagsResponse], error)
	DeleteInventoryTag(context.Context, *connect.Request[v1.DeleteInventoryTagRequest]) (*connect.Response[v1.DeleteInventoryTagResponse], error)
	CreateDataDomain(context.Context, *connect.Request[v1.CreateDataDomainRequest]) (*connect.Response[v1.CreateDataDomainResponse], error)
	ListDataDomains(context.Context, *connect.Request[v1.ListDataDomainsRequest]) (*connect.Response[v1.ListDataDomainsResponse], error)
	GetDataDomain(context.Context, *connect.Request[v1.GetDataDomainRequest]) (*connect.Response[v1.GetDataDomainResponse], error)
	UpdateDataDomain(context.Context, *connect.Request[v1.UpdateDataDomainRequest]) (*connect.Response[v1.UpdateDataDomainResponse], error)
	DeleteDataDomain(context.Context, *connect.Request[v1.DeleteDataDomainRequest]) (*connect.Response[v1.DeleteDataDomainResponse], error)
}

// NewInventoryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewInventoryServiceHandler(svc InventoryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	inventoryServiceUpdateColumnHandler := connect.NewUnaryHandler(
		InventoryServiceUpdateColumnProcedure,
		svc.UpdateColumn,
		connect.WithSchema(inventoryServiceUpdateColumnMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceCreateInventoryObjectHandler := connect.NewUnaryHandler(
		InventoryServiceCreateInventoryObjectProcedure,
		svc.CreateInventoryObject,
		connect.WithSchema(inventoryServiceCreateInventoryObjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceUpdateInventoryObjectTagsHandler := connect.NewUnaryHandler(
		InventoryServiceUpdateInventoryObjectTagsProcedure,
		svc.UpdateInventoryObjectTags,
		connect.WithSchema(inventoryServiceUpdateInventoryObjectTagsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceDeleteInventoryObjectHandler := connect.NewUnaryHandler(
		InventoryServiceDeleteInventoryObjectProcedure,
		svc.DeleteInventoryObject,
		connect.WithSchema(inventoryServiceDeleteInventoryObjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceGetInventoryObjectHandler := connect.NewUnaryHandler(
		InventoryServiceGetInventoryObjectProcedure,
		svc.GetInventoryObject,
		connect.WithSchema(inventoryServiceGetInventoryObjectMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceListInventoryObjectsHandler := connect.NewUnaryHandler(
		InventoryServiceListInventoryObjectsProcedure,
		svc.ListInventoryObjects,
		connect.WithSchema(inventoryServiceListInventoryObjectsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceCreateInventoryTagHandler := connect.NewUnaryHandler(
		InventoryServiceCreateInventoryTagProcedure,
		svc.CreateInventoryTag,
		connect.WithSchema(inventoryServiceCreateInventoryTagMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceListInventoryTagsHandler := connect.NewUnaryHandler(
		InventoryServiceListInventoryTagsProcedure,
		svc.ListInventoryTags,
		connect.WithSchema(inventoryServiceListInventoryTagsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceDeleteInventoryTagHandler := connect.NewUnaryHandler(
		InventoryServiceDeleteInventoryTagProcedure,
		svc.DeleteInventoryTag,
		connect.WithSchema(inventoryServiceDeleteInventoryTagMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceCreateDataDomainHandler := connect.NewUnaryHandler(
		InventoryServiceCreateDataDomainProcedure,
		svc.CreateDataDomain,
		connect.WithSchema(inventoryServiceCreateDataDomainMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceListDataDomainsHandler := connect.NewUnaryHandler(
		InventoryServiceListDataDomainsProcedure,
		svc.ListDataDomains,
		connect.WithSchema(inventoryServiceListDataDomainsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceGetDataDomainHandler := connect.NewUnaryHandler(
		InventoryServiceGetDataDomainProcedure,
		svc.GetDataDomain,
		connect.WithSchema(inventoryServiceGetDataDomainMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceUpdateDataDomainHandler := connect.NewUnaryHandler(
		InventoryServiceUpdateDataDomainProcedure,
		svc.UpdateDataDomain,
		connect.WithSchema(inventoryServiceUpdateDataDomainMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	inventoryServiceDeleteDataDomainHandler := connect.NewUnaryHandler(
		InventoryServiceDeleteDataDomainProcedure,
		svc.DeleteDataDomain,
		connect.WithSchema(inventoryServiceDeleteDataDomainMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/core.v1.InventoryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case InventoryServiceUpdateColumnProcedure:
			inventoryServiceUpdateColumnHandler.ServeHTTP(w, r)
		case InventoryServiceCreateInventoryObjectProcedure:
			inventoryServiceCreateInventoryObjectHandler.ServeHTTP(w, r)
		case InventoryServiceUpdateInventoryObjectTagsProcedure:
			inventoryServiceUpdateInventoryObjectTagsHandler.ServeHTTP(w, r)
		case InventoryServiceDeleteInventoryObjectProcedure:
			inventoryServiceDeleteInventoryObjectHandler.ServeHTTP(w, r)
		case InventoryServiceGetInventoryObjectProcedure:
			inventoryServiceGetInventoryObjectHandler.ServeHTTP(w, r)
		case InventoryServiceListInventoryObjectsProcedure:
			inventoryServiceListInventoryObjectsHandler.ServeHTTP(w, r)
		case InventoryServiceCreateInventoryTagProcedure:
			inventoryServiceCreateInventoryTagHandler.ServeHTTP(w, r)
		case InventoryServiceListInventoryTagsProcedure:
			inventoryServiceListInventoryTagsHandler.ServeHTTP(w, r)
		case InventoryServiceDeleteInventoryTagProcedure:
			inventoryServiceDeleteInventoryTagHandler.ServeHTTP(w, r)
		case InventoryServiceCreateDataDomainProcedure:
			inventoryServiceCreateDataDomainHandler.ServeHTTP(w, r)
		case InventoryServiceListDataDomainsProcedure:
			inventoryServiceListDataDomainsHandler.ServeHTTP(w, r)
		case InventoryServiceGetDataDomainProcedure:
			inventoryServiceGetDataDomainHandler.ServeHTTP(w, r)
		case InventoryServiceUpdateDataDomainProcedure:
			inventoryServiceUpdateDataDomainHandler.ServeHTTP(w, r)
		case InventoryServiceDeleteDataDomainProcedure:
			inventoryServiceDeleteDataDomainHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedInventoryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedInventoryServiceHandler struct{}

func (UnimplementedInventoryServiceHandler) UpdateColumn(context.Context, *connect.Request[v1.UpdateColumnRequest]) (*connect.Response[v1.UpdateColumnResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.UpdateColumn is not implemented"))
}

func (UnimplementedInventoryServiceHandler) CreateInventoryObject(context.Context, *connect.Request[v1.CreateInventoryObjectRequest]) (*connect.Response[v1.CreateInventoryObjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.CreateInventoryObject is not implemented"))
}

func (UnimplementedInventoryServiceHandler) UpdateInventoryObjectTags(context.Context, *connect.Request[v1.UpdateInventoryObjectTagsRequest]) (*connect.Response[v1.UpdateInventoryObjectTagsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.UpdateInventoryObjectTags is not implemented"))
}

func (UnimplementedInventoryServiceHandler) DeleteInventoryObject(context.Context, *connect.Request[v1.DeleteInventoryObjectRequest]) (*connect.Response[v1.DeleteInventoryObjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.DeleteInventoryObject is not implemented"))
}

func (UnimplementedInventoryServiceHandler) GetInventoryObject(context.Context, *connect.Request[v1.GetInventoryObjectRequest]) (*connect.Response[v1.GetInventoryObjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.GetInventoryObject is not implemented"))
}

func (UnimplementedInventoryServiceHandler) ListInventoryObjects(context.Context, *connect.Request[v1.ListInventoryObjectsRequest]) (*connect.Response[v1.ListInventoryObjectsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.ListInventoryObjects is not implemented"))
}

func (UnimplementedInventoryServiceHandler) CreateInventoryTag(context.Context, *connect.Request[v1.CreateInventoryTagRequest]) (*connect.Response[v1.CreateInventoryTagResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.CreateInventoryTag is not implemented"))
}

func (UnimplementedInventoryServiceHandler) ListInventoryTags(context.Context, *connect.Request[v1.ListInventoryTagsRequest]) (*connect.Response[v1.ListInventoryTagsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.ListInventoryTags is not implemented"))
}

func (UnimplementedInventoryServiceHandler) DeleteInventoryTag(context.Context, *connect.Request[v1.DeleteInventoryTagRequest]) (*connect.Response[v1.DeleteInventoryTagResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.DeleteInventoryTag is not implemented"))
}

func (UnimplementedInventoryServiceHandler) CreateDataDomain(context.Context, *connect.Request[v1.CreateDataDomainRequest]) (*connect.Response[v1.CreateDataDomainResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.CreateDataDomain is not implemented"))
}

func (UnimplementedInventoryServiceHandler) ListDataDomains(context.Context, *connect.Request[v1.ListDataDomainsRequest]) (*connect.Response[v1.ListDataDomainsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.ListDataDomains is not implemented"))
}

func (UnimplementedInventoryServiceHandler) GetDataDomain(context.Context, *connect.Request[v1.GetDataDomainRequest]) (*connect.Response[v1.GetDataDomainResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.GetDataDomain is not implemented"))
}

func (UnimplementedInventoryServiceHandler) UpdateDataDomain(context.Context, *connect.Request[v1.UpdateDataDomainRequest]) (*connect.Response[v1.UpdateDataDomainResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.UpdateDataDomain is not implemented"))
}

func (UnimplementedInventoryServiceHandler) DeleteDataDomain(context.Context, *connect.Request[v1.DeleteDataDomainRequest]) (*connect.Response[v1.DeleteDataDomainResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("core.v1.InventoryService.DeleteDataDomain is not implemented"))
}
