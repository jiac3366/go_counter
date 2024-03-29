// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: go_counter.proto

package go_counter

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
	GoCounter_Ping_FullMethodName           = "/go_counter.Go_counter/Ping"
	GoCounter_LikeNum_FullMethodName        = "/go_counter.Go_counter/LikeNum"
	GoCounter_LikeRecord_FullMethodName     = "/go_counter.Go_counter/LikeRecord"
	GoCounter_FavoriteNum_FullMethodName    = "/go_counter.Go_counter/FavoriteNum"
	GoCounter_FavoriteRecord_FullMethodName = "/go_counter.Go_counter/FavoriteRecord"
	GoCounter_ViewNum_FullMethodName        = "/go_counter.Go_counter/ViewNum"
	GoCounter_ViewRecord_FullMethodName     = "/go_counter.Go_counter/ViewRecord"
)

// GoCounterClient is the client API for GoCounter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoCounterClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// 点赞请求
	LikeNum(ctx context.Context, in *LikeNumRequest, opts ...grpc.CallOption) (*LikeNumResponse, error)
	LikeRecord(ctx context.Context, in *LikeRecordRequest, opts ...grpc.CallOption) (*LikeRecordResponse, error)
	// 收藏请求
	FavoriteNum(ctx context.Context, in *FavoriteNumRequest, opts ...grpc.CallOption) (*FavoriteNumResponse, error)
	FavoriteRecord(ctx context.Context, in *FavoriteRecordRequest, opts ...grpc.CallOption) (*FavoriteRecordResponse, error)
	// 浏览请求
	ViewNum(ctx context.Context, in *ViewNumRequest, opts ...grpc.CallOption) (*ViewNumResponse, error)
	ViewRecord(ctx context.Context, in *ViewRecordRequest, opts ...grpc.CallOption) (*ViewRecordResponse, error)
}

type goCounterClient struct {
	cc grpc.ClientConnInterface
}

func NewGoCounterClient(cc grpc.ClientConnInterface) GoCounterClient {
	return &goCounterClient{cc}
}

func (c *goCounterClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, GoCounter_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCounterClient) LikeNum(ctx context.Context, in *LikeNumRequest, opts ...grpc.CallOption) (*LikeNumResponse, error) {
	out := new(LikeNumResponse)
	err := c.cc.Invoke(ctx, GoCounter_LikeNum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCounterClient) LikeRecord(ctx context.Context, in *LikeRecordRequest, opts ...grpc.CallOption) (*LikeRecordResponse, error) {
	out := new(LikeRecordResponse)
	err := c.cc.Invoke(ctx, GoCounter_LikeRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCounterClient) FavoriteNum(ctx context.Context, in *FavoriteNumRequest, opts ...grpc.CallOption) (*FavoriteNumResponse, error) {
	out := new(FavoriteNumResponse)
	err := c.cc.Invoke(ctx, GoCounter_FavoriteNum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCounterClient) FavoriteRecord(ctx context.Context, in *FavoriteRecordRequest, opts ...grpc.CallOption) (*FavoriteRecordResponse, error) {
	out := new(FavoriteRecordResponse)
	err := c.cc.Invoke(ctx, GoCounter_FavoriteRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCounterClient) ViewNum(ctx context.Context, in *ViewNumRequest, opts ...grpc.CallOption) (*ViewNumResponse, error) {
	out := new(ViewNumResponse)
	err := c.cc.Invoke(ctx, GoCounter_ViewNum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCounterClient) ViewRecord(ctx context.Context, in *ViewRecordRequest, opts ...grpc.CallOption) (*ViewRecordResponse, error) {
	out := new(ViewRecordResponse)
	err := c.cc.Invoke(ctx, GoCounter_ViewRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoCounterServer is the server API for GoCounter service.
// All implementations must embed UnimplementedGoCounterServer
// for forward compatibility
type GoCounterServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// 点赞请求
	LikeNum(context.Context, *LikeNumRequest) (*LikeNumResponse, error)
	LikeRecord(context.Context, *LikeRecordRequest) (*LikeRecordResponse, error)
	// 收藏请求
	FavoriteNum(context.Context, *FavoriteNumRequest) (*FavoriteNumResponse, error)
	FavoriteRecord(context.Context, *FavoriteRecordRequest) (*FavoriteRecordResponse, error)
	// 浏览请求
	ViewNum(context.Context, *ViewNumRequest) (*ViewNumResponse, error)
	ViewRecord(context.Context, *ViewRecordRequest) (*ViewRecordResponse, error)
	mustEmbedUnimplementedGoCounterServer()
}

// UnimplementedGoCounterServer must be embedded to have forward compatible implementations.
type UnimplementedGoCounterServer struct {
}

func (UnimplementedGoCounterServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGoCounterServer) LikeNum(context.Context, *LikeNumRequest) (*LikeNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeNum not implemented")
}
func (UnimplementedGoCounterServer) LikeRecord(context.Context, *LikeRecordRequest) (*LikeRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeRecord not implemented")
}
func (UnimplementedGoCounterServer) FavoriteNum(context.Context, *FavoriteNumRequest) (*FavoriteNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteNum not implemented")
}
func (UnimplementedGoCounterServer) FavoriteRecord(context.Context, *FavoriteRecordRequest) (*FavoriteRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteRecord not implemented")
}
func (UnimplementedGoCounterServer) ViewNum(context.Context, *ViewNumRequest) (*ViewNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewNum not implemented")
}
func (UnimplementedGoCounterServer) ViewRecord(context.Context, *ViewRecordRequest) (*ViewRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewRecord not implemented")
}
func (UnimplementedGoCounterServer) mustEmbedUnimplementedGoCounterServer() {}

// UnsafeGoCounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoCounterServer will
// result in compilation errors.
type UnsafeGoCounterServer interface {
	mustEmbedUnimplementedGoCounterServer()
}

func RegisterGoCounterServer(s grpc.ServiceRegistrar, srv GoCounterServer) {
	s.RegisterService(&GoCounter_ServiceDesc, srv)
}

func _GoCounter_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCounterServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoCounter_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCounterServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCounter_LikeNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCounterServer).LikeNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoCounter_LikeNum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCounterServer).LikeNum(ctx, req.(*LikeNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCounter_LikeRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCounterServer).LikeRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoCounter_LikeRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCounterServer).LikeRecord(ctx, req.(*LikeRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCounter_FavoriteNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCounterServer).FavoriteNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoCounter_FavoriteNum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCounterServer).FavoriteNum(ctx, req.(*FavoriteNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCounter_FavoriteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCounterServer).FavoriteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoCounter_FavoriteRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCounterServer).FavoriteRecord(ctx, req.(*FavoriteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCounter_ViewNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCounterServer).ViewNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoCounter_ViewNum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCounterServer).ViewNum(ctx, req.(*ViewNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCounter_ViewRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCounterServer).ViewRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoCounter_ViewRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCounterServer).ViewRecord(ctx, req.(*ViewRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoCounter_ServiceDesc is the grpc.ServiceDesc for GoCounter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoCounter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_counter.Go_counter",
	HandlerType: (*GoCounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GoCounter_Ping_Handler,
		},
		{
			MethodName: "LikeNum",
			Handler:    _GoCounter_LikeNum_Handler,
		},
		{
			MethodName: "LikeRecord",
			Handler:    _GoCounter_LikeRecord_Handler,
		},
		{
			MethodName: "FavoriteNum",
			Handler:    _GoCounter_FavoriteNum_Handler,
		},
		{
			MethodName: "FavoriteRecord",
			Handler:    _GoCounter_FavoriteRecord_Handler,
		},
		{
			MethodName: "ViewNum",
			Handler:    _GoCounter_ViewNum_Handler,
		},
		{
			MethodName: "ViewRecord",
			Handler:    _GoCounter_ViewRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go_counter.proto",
}
