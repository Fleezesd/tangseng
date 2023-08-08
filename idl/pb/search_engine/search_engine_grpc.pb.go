// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: search_engine.proto

package search_engine

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchEngineServiceClient is the client API for SearchEngineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchEngineServiceClient interface {
	SearchEngineSearch(ctx context.Context, in *SearchEngineRequest, opts ...grpc.CallOption) (*SearchEngineResponse, error)
}

type searchEngineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchEngineServiceClient(cc grpc.ClientConnInterface) SearchEngineServiceClient {
	return &searchEngineServiceClient{cc}
}

func (c *searchEngineServiceClient) SearchEngineSearch(ctx context.Context, in *SearchEngineRequest, opts ...grpc.CallOption) (*SearchEngineResponse, error) {
	out := new(SearchEngineResponse)
	err := c.cc.Invoke(ctx, "/SearchEngineService/SearchEngineSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchEngineServiceServer is the server API for SearchEngineService service.
// All implementations must embed UnimplementedSearchEngineServiceServer
// for forward compatibility
type SearchEngineServiceServer interface {
	SearchEngineSearch(context.Context, *SearchEngineRequest) (*SearchEngineResponse, error)
	mustEmbedUnimplementedSearchEngineServiceServer()
}

// UnimplementedSearchEngineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchEngineServiceServer struct {
}

func (UnimplementedSearchEngineServiceServer) SearchEngineSearch(context.Context, *SearchEngineRequest) (*SearchEngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEngineSearch not implemented")
}
func (UnimplementedSearchEngineServiceServer) mustEmbedUnimplementedSearchEngineServiceServer() {}

// UnsafeSearchEngineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchEngineServiceServer will
// result in compilation errors.
type UnsafeSearchEngineServiceServer interface {
	mustEmbedUnimplementedSearchEngineServiceServer()
}

func RegisterSearchEngineServiceServer(s grpc.ServiceRegistrar, srv SearchEngineServiceServer) {
	s.RegisterService(&SearchEngineService_ServiceDesc, srv)
}

func _SearchEngineService_SearchEngineSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchEngineServiceServer).SearchEngineSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchEngineService/SearchEngineSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchEngineServiceServer).SearchEngineSearch(ctx, req.(*SearchEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchEngineService_ServiceDesc is the grpc.ServiceDesc for SearchEngineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchEngineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SearchEngineService",
	HandlerType: (*SearchEngineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchEngineSearch",
			Handler:    _SearchEngineService_SearchEngineSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search_engine.proto",
}
