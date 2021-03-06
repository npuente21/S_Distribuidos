// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_usermsg_grpc

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

// LiderServicesClient is the client API for LiderServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiderServicesClient interface {
	NewPlayer(ctx context.Context, in *Message, opts ...grpc.CallOption) (*User, error)
	Luz_Roja_Verde(ctx context.Context, in *Jugada_1, opts ...grpc.CallOption) (*Resp_1, error)
	Pozo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Monto, error)
}

type liderServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewLiderServicesClient(cc grpc.ClientConnInterface) LiderServicesClient {
	return &liderServicesClient{cc}
}

func (c *liderServicesClient) NewPlayer(ctx context.Context, in *Message, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/usermsg.LiderServices/NewPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServicesClient) Luz_Roja_Verde(ctx context.Context, in *Jugada_1, opts ...grpc.CallOption) (*Resp_1, error) {
	out := new(Resp_1)
	err := c.cc.Invoke(ctx, "/usermsg.LiderServices/Luz_Roja_Verde", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServicesClient) Pozo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Monto, error) {
	out := new(Monto)
	err := c.cc.Invoke(ctx, "/usermsg.LiderServices/Pozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiderServicesServer is the server API for LiderServices service.
// All implementations must embed UnimplementedLiderServicesServer
// for forward compatibility
type LiderServicesServer interface {
	NewPlayer(context.Context, *Message) (*User, error)
	Luz_Roja_Verde(context.Context, *Jugada_1) (*Resp_1, error)
	Pozo(context.Context, *Req) (*Monto, error)
	mustEmbedUnimplementedLiderServicesServer()
}

// UnimplementedLiderServicesServer must be embedded to have forward compatible implementations.
type UnimplementedLiderServicesServer struct {
}

func (UnimplementedLiderServicesServer) NewPlayer(context.Context, *Message) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewPlayer not implemented")
}
func (UnimplementedLiderServicesServer) Luz_Roja_Verde(context.Context, *Jugada_1) (*Resp_1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Luz_Roja_Verde not implemented")
}
func (UnimplementedLiderServicesServer) Pozo(context.Context, *Req) (*Monto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pozo not implemented")
}
func (UnimplementedLiderServicesServer) mustEmbedUnimplementedLiderServicesServer() {}

// UnsafeLiderServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiderServicesServer will
// result in compilation errors.
type UnsafeLiderServicesServer interface {
	mustEmbedUnimplementedLiderServicesServer()
}

func RegisterLiderServicesServer(s grpc.ServiceRegistrar, srv LiderServicesServer) {
	s.RegisterService(&LiderServices_ServiceDesc, srv)
}

func _LiderServices_NewPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServicesServer).NewPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.LiderServices/NewPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServicesServer).NewPlayer(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderServices_Luz_Roja_Verde_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada_1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServicesServer).Luz_Roja_Verde(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.LiderServices/Luz_Roja_Verde",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServicesServer).Luz_Roja_Verde(ctx, req.(*Jugada_1))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderServices_Pozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServicesServer).Pozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.LiderServices/Pozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServicesServer).Pozo(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// LiderServices_ServiceDesc is the grpc.ServiceDesc for LiderServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiderServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermsg.LiderServices",
	HandlerType: (*LiderServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPlayer",
			Handler:    _LiderServices_NewPlayer_Handler,
		},
		{
			MethodName: "Luz_Roja_Verde",
			Handler:    _LiderServices_Luz_Roja_Verde_Handler,
		},
		{
			MethodName: "Pozo",
			Handler:    _LiderServices_Pozo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermsg/usermsg.proto",
}

// PozoServicesClient is the client API for PozoServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PozoServicesClient interface {
	MontoPozo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Monto, error)
}

type pozoServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewPozoServicesClient(cc grpc.ClientConnInterface) PozoServicesClient {
	return &pozoServicesClient{cc}
}

func (c *pozoServicesClient) MontoPozo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Monto, error) {
	out := new(Monto)
	err := c.cc.Invoke(ctx, "/usermsg.PozoServices/MontoPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PozoServicesServer is the server API for PozoServices service.
// All implementations must embed UnimplementedPozoServicesServer
// for forward compatibility
type PozoServicesServer interface {
	MontoPozo(context.Context, *Req) (*Monto, error)
	mustEmbedUnimplementedPozoServicesServer()
}

// UnimplementedPozoServicesServer must be embedded to have forward compatible implementations.
type UnimplementedPozoServicesServer struct {
}

func (UnimplementedPozoServicesServer) MontoPozo(context.Context, *Req) (*Monto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MontoPozo not implemented")
}
func (UnimplementedPozoServicesServer) mustEmbedUnimplementedPozoServicesServer() {}

// UnsafePozoServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PozoServicesServer will
// result in compilation errors.
type UnsafePozoServicesServer interface {
	mustEmbedUnimplementedPozoServicesServer()
}

func RegisterPozoServicesServer(s grpc.ServiceRegistrar, srv PozoServicesServer) {
	s.RegisterService(&PozoServices_ServiceDesc, srv)
}

func _PozoServices_MontoPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PozoServicesServer).MontoPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.PozoServices/MontoPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PozoServicesServer).MontoPozo(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// PozoServices_ServiceDesc is the grpc.ServiceDesc for PozoServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PozoServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermsg.PozoServices",
	HandlerType: (*PozoServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MontoPozo",
			Handler:    _PozoServices_MontoPozo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermsg/usermsg.proto",
}

// NameNodeClient is the client API for NameNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NameNodeClient interface {
	JugadaPlayer(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Registro, error)
}

type nameNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeClient(cc grpc.ClientConnInterface) NameNodeClient {
	return &nameNodeClient{cc}
}

func (c *nameNodeClient) JugadaPlayer(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Registro, error) {
	out := new(Registro)
	err := c.cc.Invoke(ctx, "/usermsg.NameNode/JugadaPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameNodeServer is the server API for NameNode service.
// All implementations must embed UnimplementedNameNodeServer
// for forward compatibility
type NameNodeServer interface {
	JugadaPlayer(context.Context, *Jugada) (*Registro, error)
	mustEmbedUnimplementedNameNodeServer()
}

// UnimplementedNameNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNameNodeServer struct {
}

func (UnimplementedNameNodeServer) JugadaPlayer(context.Context, *Jugada) (*Registro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JugadaPlayer not implemented")
}
func (UnimplementedNameNodeServer) mustEmbedUnimplementedNameNodeServer() {}

// UnsafeNameNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NameNodeServer will
// result in compilation errors.
type UnsafeNameNodeServer interface {
	mustEmbedUnimplementedNameNodeServer()
}

func RegisterNameNodeServer(s grpc.ServiceRegistrar, srv NameNodeServer) {
	s.RegisterService(&NameNode_ServiceDesc, srv)
}

func _NameNode_JugadaPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServer).JugadaPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.NameNode/JugadaPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServer).JugadaPlayer(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

// NameNode_ServiceDesc is the grpc.ServiceDesc for NameNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NameNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermsg.NameNode",
	HandlerType: (*NameNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JugadaPlayer",
			Handler:    _NameNode_JugadaPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermsg/usermsg.proto",
}

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeClient interface {
	RegistrarInfo(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Check, error)
}

type dataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeClient(cc grpc.ClientConnInterface) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) RegistrarInfo(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Check, error) {
	out := new(Check)
	err := c.cc.Invoke(ctx, "/usermsg.DataNode/RegistrarInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServer is the server API for DataNode service.
// All implementations must embed UnimplementedDataNodeServer
// for forward compatibility
type DataNodeServer interface {
	RegistrarInfo(context.Context, *Jugada) (*Check, error)
	mustEmbedUnimplementedDataNodeServer()
}

// UnimplementedDataNodeServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (UnimplementedDataNodeServer) RegistrarInfo(context.Context, *Jugada) (*Check, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrarInfo not implemented")
}
func (UnimplementedDataNodeServer) mustEmbedUnimplementedDataNodeServer() {}

// UnsafeDataNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServer will
// result in compilation errors.
type UnsafeDataNodeServer interface {
	mustEmbedUnimplementedDataNodeServer()
}

func RegisterDataNodeServer(s grpc.ServiceRegistrar, srv DataNodeServer) {
	s.RegisterService(&DataNode_ServiceDesc, srv)
}

func _DataNode_RegistrarInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).RegistrarInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermsg.DataNode/RegistrarInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).RegistrarInfo(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNode_ServiceDesc is the grpc.ServiceDesc for DataNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermsg.DataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistrarInfo",
			Handler:    _DataNode_RegistrarInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usermsg/usermsg.proto",
}
