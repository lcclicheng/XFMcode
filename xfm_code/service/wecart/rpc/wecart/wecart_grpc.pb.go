// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: wecart.proto

package wecart

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
	Wecart_PlaceAnOrderLogic_FullMethodName = "/wecart.Wecart/PlaceAnOrderLogic"
)

// WecartClient is the client API for Wecart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WecartClient interface {
	PlaceAnOrderLogic(ctx context.Context, in *PlaceAnOrderReq, opts ...grpc.CallOption) (*PlaceAnOrderResp, error)
}

type wecartClient struct {
	cc grpc.ClientConnInterface
}

func NewWecartClient(cc grpc.ClientConnInterface) WecartClient {
	return &wecartClient{cc}
}

func (c *wecartClient) PlaceAnOrderLogic(ctx context.Context, in *PlaceAnOrderReq, opts ...grpc.CallOption) (*PlaceAnOrderResp, error) {
	out := new(PlaceAnOrderResp)
	err := c.cc.Invoke(ctx, Wecart_PlaceAnOrderLogic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WecartServer is the server API for Wecart service.
// All implementations must embed UnimplementedWecartServer
// for forward compatibility
type WecartServer interface {
	PlaceAnOrderLogic(context.Context, *PlaceAnOrderReq) (*PlaceAnOrderResp, error)
	mustEmbedUnimplementedWecartServer()
}

// UnimplementedWecartServer must be embedded to have forward compatible implementations.
type UnimplementedWecartServer struct {
}

func (UnimplementedWecartServer) PlaceAnOrderLogic(context.Context, *PlaceAnOrderReq) (*PlaceAnOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceAnOrderLogic not implemented")
}
func (UnimplementedWecartServer) mustEmbedUnimplementedWecartServer() {}

// UnsafeWecartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WecartServer will
// result in compilation errors.
type UnsafeWecartServer interface {
	mustEmbedUnimplementedWecartServer()
}

func RegisterWecartServer(s grpc.ServiceRegistrar, srv WecartServer) {
	s.RegisterService(&Wecart_ServiceDesc, srv)
}

func _Wecart_PlaceAnOrderLogic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceAnOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WecartServer).PlaceAnOrderLogic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wecart_PlaceAnOrderLogic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WecartServer).PlaceAnOrderLogic(ctx, req.(*PlaceAnOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Wecart_ServiceDesc is the grpc.ServiceDesc for Wecart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wecart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wecart.Wecart",
	HandlerType: (*WecartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceAnOrderLogic",
			Handler:    _Wecart_PlaceAnOrderLogic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wecart.proto",
}
