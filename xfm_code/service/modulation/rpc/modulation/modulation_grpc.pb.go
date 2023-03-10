// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.1
// source: modulation.proto

package modulation

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

// ModulationRpcClient is the client API for ModulationRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModulationRpcClient interface {
	//查询消费码状态
	QueryCodeStatus(ctx context.Context, in *CodeStatusRequest, opts ...grpc.CallOption) (*CodeStatusResponse, error)
	//用户申请消费码返回给用户
	RequestConsumption(ctx context.Context, in *RequestConsumptionRequest, opts ...grpc.CallOption) (*RequestConsumptionResponse, error)
	// 二维码订单查询
	OrderDetailsLogic(ctx context.Context, in *OrderDetailsReq, opts ...grpc.CallOption) (*OrderDetailsResp, error)
}

type modulationRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewModulationRpcClient(cc grpc.ClientConnInterface) ModulationRpcClient {
	return &modulationRpcClient{cc}
}

func (c *modulationRpcClient) QueryCodeStatus(ctx context.Context, in *CodeStatusRequest, opts ...grpc.CallOption) (*CodeStatusResponse, error) {
	out := new(CodeStatusResponse)
	err := c.cc.Invoke(ctx, "/modulation.modulationRpc/QueryCodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modulationRpcClient) RequestConsumption(ctx context.Context, in *RequestConsumptionRequest, opts ...grpc.CallOption) (*RequestConsumptionResponse, error) {
	out := new(RequestConsumptionResponse)
	err := c.cc.Invoke(ctx, "/modulation.modulationRpc/RequestConsumption", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modulationRpcClient) OrderDetailsLogic(ctx context.Context, in *OrderDetailsReq, opts ...grpc.CallOption) (*OrderDetailsResp, error) {
	out := new(OrderDetailsResp)
	err := c.cc.Invoke(ctx, "/modulation.modulationRpc/OrderDetailsLogic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModulationRpcServer is the server API for ModulationRpc service.
// All implementations must embed UnimplementedModulationRpcServer
// for forward compatibility
type ModulationRpcServer interface {
	//查询消费码状态
	QueryCodeStatus(context.Context, *CodeStatusRequest) (*CodeStatusResponse, error)
	//用户申请消费码返回给用户
	RequestConsumption(context.Context, *RequestConsumptionRequest) (*RequestConsumptionResponse, error)
	// 二维码订单查询
	OrderDetailsLogic(context.Context, *OrderDetailsReq) (*OrderDetailsResp, error)
	mustEmbedUnimplementedModulationRpcServer()
}

// UnimplementedModulationRpcServer must be embedded to have forward compatible implementations.
type UnimplementedModulationRpcServer struct {
}

func (UnimplementedModulationRpcServer) QueryCodeStatus(context.Context, *CodeStatusRequest) (*CodeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCodeStatus not implemented")
}
func (UnimplementedModulationRpcServer) RequestConsumption(context.Context, *RequestConsumptionRequest) (*RequestConsumptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestConsumption not implemented")
}
func (UnimplementedModulationRpcServer) OrderDetailsLogic(context.Context, *OrderDetailsReq) (*OrderDetailsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetailsLogic not implemented")
}
func (UnimplementedModulationRpcServer) mustEmbedUnimplementedModulationRpcServer() {}

// UnsafeModulationRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModulationRpcServer will
// result in compilation errors.
type UnsafeModulationRpcServer interface {
	mustEmbedUnimplementedModulationRpcServer()
}

func RegisterModulationRpcServer(s grpc.ServiceRegistrar, srv ModulationRpcServer) {
	s.RegisterService(&ModulationRpc_ServiceDesc, srv)
}

func _ModulationRpc_QueryCodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModulationRpcServer).QueryCodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modulation.modulationRpc/QueryCodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModulationRpcServer).QueryCodeStatus(ctx, req.(*CodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModulationRpc_RequestConsumption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestConsumptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModulationRpcServer).RequestConsumption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modulation.modulationRpc/RequestConsumption",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModulationRpcServer).RequestConsumption(ctx, req.(*RequestConsumptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModulationRpc_OrderDetailsLogic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDetailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModulationRpcServer).OrderDetailsLogic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modulation.modulationRpc/OrderDetailsLogic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModulationRpcServer).OrderDetailsLogic(ctx, req.(*OrderDetailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ModulationRpc_ServiceDesc is the grpc.ServiceDesc for ModulationRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModulationRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "modulation.modulationRpc",
	HandlerType: (*ModulationRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryCodeStatus",
			Handler:    _ModulationRpc_QueryCodeStatus_Handler,
		},
		{
			MethodName: "RequestConsumption",
			Handler:    _ModulationRpc_RequestConsumption_Handler,
		},
		{
			MethodName: "OrderDetailsLogic",
			Handler:    _ModulationRpc_OrderDetailsLogic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modulation.proto",
}
