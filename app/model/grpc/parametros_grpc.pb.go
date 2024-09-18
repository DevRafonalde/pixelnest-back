// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: parametros.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Parametros_FindAllParametros_FullMethodName   = "/grpc.Parametros/FindAllParametros"
	Parametros_FindParametroByNome_FullMethodName = "/grpc.Parametros/FindParametroByNome"
	Parametros_FindParametroById_FullMethodName   = "/grpc.Parametros/FindParametroById"
	Parametros_CreateParametro_FullMethodName     = "/grpc.Parametros/CreateParametro"
	Parametros_UpdateParametro_FullMethodName     = "/grpc.Parametros/UpdateParametro"
	Parametros_DeleteParametro_FullMethodName     = "/grpc.Parametros/DeleteParametro"
)

// ParametrosClient is the client API for Parametros service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Definição do serviço
type ParametrosClient interface {
	FindAllParametros(ctx context.Context, in *RequestVazio, opts ...grpc.CallOption) (*ListaParametros, error)
	FindParametroByNome(ctx context.Context, in *RequestNome, opts ...grpc.CallOption) (*Parametro, error)
	FindParametroById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Parametro, error)
	CreateParametro(ctx context.Context, in *Parametro, opts ...grpc.CallOption) (*Parametro, error)
	UpdateParametro(ctx context.Context, in *Parametro, opts ...grpc.CallOption) (*Parametro, error)
	DeleteParametro(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error)
}

type parametrosClient struct {
	cc grpc.ClientConnInterface
}

func NewParametrosClient(cc grpc.ClientConnInterface) ParametrosClient {
	return &parametrosClient{cc}
}

func (c *parametrosClient) FindAllParametros(ctx context.Context, in *RequestVazio, opts ...grpc.CallOption) (*ListaParametros, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaParametros)
	err := c.cc.Invoke(ctx, Parametros_FindAllParametros_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parametrosClient) FindParametroByNome(ctx context.Context, in *RequestNome, opts ...grpc.CallOption) (*Parametro, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Parametro)
	err := c.cc.Invoke(ctx, Parametros_FindParametroByNome_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parametrosClient) FindParametroById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Parametro, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Parametro)
	err := c.cc.Invoke(ctx, Parametros_FindParametroById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parametrosClient) CreateParametro(ctx context.Context, in *Parametro, opts ...grpc.CallOption) (*Parametro, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Parametro)
	err := c.cc.Invoke(ctx, Parametros_CreateParametro_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parametrosClient) UpdateParametro(ctx context.Context, in *Parametro, opts ...grpc.CallOption) (*Parametro, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Parametro)
	err := c.cc.Invoke(ctx, Parametros_UpdateParametro_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parametrosClient) DeleteParametro(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseBool)
	err := c.cc.Invoke(ctx, Parametros_DeleteParametro_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParametrosServer is the server API for Parametros service.
// All implementations must embed UnimplementedParametrosServer
// for forward compatibility.
//
// Definição do serviço
type ParametrosServer interface {
	FindAllParametros(context.Context, *RequestVazio) (*ListaParametros, error)
	FindParametroByNome(context.Context, *RequestNome) (*Parametro, error)
	FindParametroById(context.Context, *RequestId) (*Parametro, error)
	CreateParametro(context.Context, *Parametro) (*Parametro, error)
	UpdateParametro(context.Context, *Parametro) (*Parametro, error)
	DeleteParametro(context.Context, *RequestId) (*ResponseBool, error)
	mustEmbedUnimplementedParametrosServer()
}

// UnimplementedParametrosServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParametrosServer struct{}

func (UnimplementedParametrosServer) FindAllParametros(context.Context, *RequestVazio) (*ListaParametros, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllParametros not implemented")
}
func (UnimplementedParametrosServer) FindParametroByNome(context.Context, *RequestNome) (*Parametro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindParametroByNome not implemented")
}
func (UnimplementedParametrosServer) FindParametroById(context.Context, *RequestId) (*Parametro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindParametroById not implemented")
}
func (UnimplementedParametrosServer) CreateParametro(context.Context, *Parametro) (*Parametro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParametro not implemented")
}
func (UnimplementedParametrosServer) UpdateParametro(context.Context, *Parametro) (*Parametro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParametro not implemented")
}
func (UnimplementedParametrosServer) DeleteParametro(context.Context, *RequestId) (*ResponseBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParametro not implemented")
}
func (UnimplementedParametrosServer) mustEmbedUnimplementedParametrosServer() {}
func (UnimplementedParametrosServer) testEmbeddedByValue()                    {}

// UnsafeParametrosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParametrosServer will
// result in compilation errors.
type UnsafeParametrosServer interface {
	mustEmbedUnimplementedParametrosServer()
}

func RegisterParametrosServer(s grpc.ServiceRegistrar, srv ParametrosServer) {
	// If the following call pancis, it indicates UnimplementedParametrosServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Parametros_ServiceDesc, srv)
}

func _Parametros_FindAllParametros_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVazio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParametrosServer).FindAllParametros(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Parametros_FindAllParametros_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParametrosServer).FindAllParametros(ctx, req.(*RequestVazio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parametros_FindParametroByNome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNome)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParametrosServer).FindParametroByNome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Parametros_FindParametroByNome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParametrosServer).FindParametroByNome(ctx, req.(*RequestNome))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parametros_FindParametroById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParametrosServer).FindParametroById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Parametros_FindParametroById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParametrosServer).FindParametroById(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parametros_CreateParametro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parametro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParametrosServer).CreateParametro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Parametros_CreateParametro_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParametrosServer).CreateParametro(ctx, req.(*Parametro))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parametros_UpdateParametro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Parametro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParametrosServer).UpdateParametro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Parametros_UpdateParametro_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParametrosServer).UpdateParametro(ctx, req.(*Parametro))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parametros_DeleteParametro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParametrosServer).DeleteParametro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Parametros_DeleteParametro_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParametrosServer).DeleteParametro(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// Parametros_ServiceDesc is the grpc.ServiceDesc for Parametros service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parametros_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Parametros",
	HandlerType: (*ParametrosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllParametros",
			Handler:    _Parametros_FindAllParametros_Handler,
		},
		{
			MethodName: "FindParametroByNome",
			Handler:    _Parametros_FindParametroByNome_Handler,
		},
		{
			MethodName: "FindParametroById",
			Handler:    _Parametros_FindParametroById_Handler,
		},
		{
			MethodName: "CreateParametro",
			Handler:    _Parametros_CreateParametro_Handler,
		},
		{
			MethodName: "UpdateParametro",
			Handler:    _Parametros_UpdateParametro_Handler,
		},
		{
			MethodName: "DeleteParametro",
			Handler:    _Parametros_DeleteParametro_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parametros.proto",
}
