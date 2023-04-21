// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: warehouse.proto

package proto

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
	WarehouseManagementService_RenderWarehouse_FullMethodName = "/warehouse.WarehouseManagementService/RenderWarehouse"
	WarehouseManagementService_InfoWarehouse_FullMethodName   = "/warehouse.WarehouseManagementService/InfoWarehouse"
	WarehouseManagementService_AddStock_FullMethodName        = "/warehouse.WarehouseManagementService/AddStock"
)

// WarehouseManagementServiceClient is the client API for WarehouseManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarehouseManagementServiceClient interface {
	RenderWarehouse(ctx context.Context, in *RenderWarehouseRequest, opts ...grpc.CallOption) (*RenderWarehouseResponse, error)
	InfoWarehouse(ctx context.Context, in *InfoWarehouseRequest, opts ...grpc.CallOption) (*InfoWarehouseResponse, error)
	AddStock(ctx context.Context, in *AddStockRequest, opts ...grpc.CallOption) (*AddStockResponse, error)
}

type warehouseManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarehouseManagementServiceClient(cc grpc.ClientConnInterface) WarehouseManagementServiceClient {
	return &warehouseManagementServiceClient{cc}
}

func (c *warehouseManagementServiceClient) RenderWarehouse(ctx context.Context, in *RenderWarehouseRequest, opts ...grpc.CallOption) (*RenderWarehouseResponse, error) {
	out := new(RenderWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseManagementService_RenderWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseManagementServiceClient) InfoWarehouse(ctx context.Context, in *InfoWarehouseRequest, opts ...grpc.CallOption) (*InfoWarehouseResponse, error) {
	out := new(InfoWarehouseResponse)
	err := c.cc.Invoke(ctx, WarehouseManagementService_InfoWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseManagementServiceClient) AddStock(ctx context.Context, in *AddStockRequest, opts ...grpc.CallOption) (*AddStockResponse, error) {
	out := new(AddStockResponse)
	err := c.cc.Invoke(ctx, WarehouseManagementService_AddStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseManagementServiceServer is the server API for WarehouseManagementService service.
// All implementations must embed UnimplementedWarehouseManagementServiceServer
// for forward compatibility
type WarehouseManagementServiceServer interface {
	RenderWarehouse(context.Context, *RenderWarehouseRequest) (*RenderWarehouseResponse, error)
	InfoWarehouse(context.Context, *InfoWarehouseRequest) (*InfoWarehouseResponse, error)
	AddStock(context.Context, *AddStockRequest) (*AddStockResponse, error)
	mustEmbedUnimplementedWarehouseManagementServiceServer()
}

// UnimplementedWarehouseManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWarehouseManagementServiceServer struct {
}

func (UnimplementedWarehouseManagementServiceServer) RenderWarehouse(context.Context, *RenderWarehouseRequest) (*RenderWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderWarehouse not implemented")
}
func (UnimplementedWarehouseManagementServiceServer) InfoWarehouse(context.Context, *InfoWarehouseRequest) (*InfoWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoWarehouse not implemented")
}
func (UnimplementedWarehouseManagementServiceServer) AddStock(context.Context, *AddStockRequest) (*AddStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStock not implemented")
}
func (UnimplementedWarehouseManagementServiceServer) mustEmbedUnimplementedWarehouseManagementServiceServer() {
}

// UnsafeWarehouseManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarehouseManagementServiceServer will
// result in compilation errors.
type UnsafeWarehouseManagementServiceServer interface {
	mustEmbedUnimplementedWarehouseManagementServiceServer()
}

func RegisterWarehouseManagementServiceServer(s grpc.ServiceRegistrar, srv WarehouseManagementServiceServer) {
	s.RegisterService(&WarehouseManagementService_ServiceDesc, srv)
}

func _WarehouseManagementService_RenderWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseManagementServiceServer).RenderWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseManagementService_RenderWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseManagementServiceServer).RenderWarehouse(ctx, req.(*RenderWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseManagementService_InfoWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseManagementServiceServer).InfoWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseManagementService_InfoWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseManagementServiceServer).InfoWarehouse(ctx, req.(*InfoWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarehouseManagementService_AddStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseManagementServiceServer).AddStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarehouseManagementService_AddStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseManagementServiceServer).AddStock(ctx, req.(*AddStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WarehouseManagementService_ServiceDesc is the grpc.ServiceDesc for WarehouseManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WarehouseManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warehouse.WarehouseManagementService",
	HandlerType: (*WarehouseManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RenderWarehouse",
			Handler:    _WarehouseManagementService_RenderWarehouse_Handler,
		},
		{
			MethodName: "InfoWarehouse",
			Handler:    _WarehouseManagementService_InfoWarehouse_Handler,
		},
		{
			MethodName: "AddStock",
			Handler:    _WarehouseManagementService_AddStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse.proto",
}
