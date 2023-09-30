// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: dbservice.proto

package dbservice

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

const (
	DatabaseService_AddGame_FullMethodName = "/db.DatabaseService/AddGame"
	DatabaseService_GetGame_FullMethodName = "/db.DatabaseService/GetGame"
)

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	AddGame(ctx context.Context, in *AddGameRequest, opts ...grpc.CallOption) (*AddGameResponse, error)
	GetGame(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*GetGameResponse, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) AddGame(ctx context.Context, in *AddGameRequest, opts ...grpc.CallOption) (*AddGameResponse, error) {
	out := new(AddGameResponse)
	err := c.cc.Invoke(ctx, DatabaseService_AddGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetGame(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*GetGameResponse, error) {
	out := new(GetGameResponse)
	err := c.cc.Invoke(ctx, DatabaseService_GetGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations must embed UnimplementedDatabaseServiceServer
// for forward compatibility
type DatabaseServiceServer interface {
	AddGame(context.Context, *AddGameRequest) (*AddGameResponse, error)
	GetGame(context.Context, *GetGameRequest) (*GetGameResponse, error)
	mustEmbedUnimplementedDatabaseServiceServer()
}

// UnimplementedDatabaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (UnimplementedDatabaseServiceServer) AddGame(context.Context, *AddGameRequest) (*AddGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGame not implemented")
}
func (UnimplementedDatabaseServiceServer) GetGame(context.Context, *GetGameRequest) (*GetGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedDatabaseServiceServer) mustEmbedUnimplementedDatabaseServiceServer() {}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_AddGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).AddGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_AddGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).AddGame(ctx, req.(*AddGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_GetGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetGame(ctx, req.(*GetGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGame",
			Handler:    _DatabaseService_AddGame_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _DatabaseService_GetGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbservice.proto",
}