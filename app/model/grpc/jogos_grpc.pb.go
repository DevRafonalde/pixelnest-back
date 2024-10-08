// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: jogos.proto

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
	Jogos_FindAllJogos_FullMethodName              = "/grpc.Jogos/FindAllJogos"
	Jogos_FindJogoById_FullMethodName              = "/grpc.Jogos/FindJogoById"
	Jogos_FindJogoByNome_FullMethodName            = "/grpc.Jogos/FindJogoByNome"
	Jogos_FindJogoByGenero_FullMethodName          = "/grpc.Jogos/FindJogoByGenero"
	Jogos_FindJogoByUsuario_FullMethodName         = "/grpc.Jogos/FindJogoByUsuario"
	Jogos_FindJogoFavoritoByUsuario_FullMethodName = "/grpc.Jogos/FindJogoFavoritoByUsuario"
	Jogos_CreateJogo_FullMethodName                = "/grpc.Jogos/CreateJogo"
	Jogos_UpdateJogo_FullMethodName                = "/grpc.Jogos/UpdateJogo"
	Jogos_DeleteJogo_FullMethodName                = "/grpc.Jogos/DeleteJogo"
)

// JogosClient is the client API for Jogos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Definição do serviço
type JogosClient interface {
	FindAllJogos(ctx context.Context, in *RequestVazio, opts ...grpc.CallOption) (*ListaJogos, error)
	FindJogoById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Jogo, error)
	FindJogoByNome(ctx context.Context, in *RequestNome, opts ...grpc.CallOption) (*ListaJogos, error)
	FindJogoByGenero(ctx context.Context, in *RequestNome, opts ...grpc.CallOption) (*ListaJogos, error)
	FindJogoByUsuario(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ListaJogos, error)
	FindJogoFavoritoByUsuario(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ListaJogos, error)
	CreateJogo(ctx context.Context, in *Jogo, opts ...grpc.CallOption) (*Jogo, error)
	UpdateJogo(ctx context.Context, in *Jogo, opts ...grpc.CallOption) (*Jogo, error)
	DeleteJogo(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error)
}

type jogosClient struct {
	cc grpc.ClientConnInterface
}

func NewJogosClient(cc grpc.ClientConnInterface) JogosClient {
	return &jogosClient{cc}
}

func (c *jogosClient) FindAllJogos(ctx context.Context, in *RequestVazio, opts ...grpc.CallOption) (*ListaJogos, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaJogos)
	err := c.cc.Invoke(ctx, Jogos_FindAllJogos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) FindJogoById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*Jogo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Jogo)
	err := c.cc.Invoke(ctx, Jogos_FindJogoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) FindJogoByNome(ctx context.Context, in *RequestNome, opts ...grpc.CallOption) (*ListaJogos, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaJogos)
	err := c.cc.Invoke(ctx, Jogos_FindJogoByNome_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) FindJogoByGenero(ctx context.Context, in *RequestNome, opts ...grpc.CallOption) (*ListaJogos, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaJogos)
	err := c.cc.Invoke(ctx, Jogos_FindJogoByGenero_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) FindJogoByUsuario(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ListaJogos, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaJogos)
	err := c.cc.Invoke(ctx, Jogos_FindJogoByUsuario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) FindJogoFavoritoByUsuario(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ListaJogos, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaJogos)
	err := c.cc.Invoke(ctx, Jogos_FindJogoFavoritoByUsuario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) CreateJogo(ctx context.Context, in *Jogo, opts ...grpc.CallOption) (*Jogo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Jogo)
	err := c.cc.Invoke(ctx, Jogos_CreateJogo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) UpdateJogo(ctx context.Context, in *Jogo, opts ...grpc.CallOption) (*Jogo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Jogo)
	err := c.cc.Invoke(ctx, Jogos_UpdateJogo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jogosClient) DeleteJogo(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseBool)
	err := c.cc.Invoke(ctx, Jogos_DeleteJogo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JogosServer is the server API for Jogos service.
// All implementations must embed UnimplementedJogosServer
// for forward compatibility.
//
// Definição do serviço
type JogosServer interface {
	FindAllJogos(context.Context, *RequestVazio) (*ListaJogos, error)
	FindJogoById(context.Context, *RequestId) (*Jogo, error)
	FindJogoByNome(context.Context, *RequestNome) (*ListaJogos, error)
	FindJogoByGenero(context.Context, *RequestNome) (*ListaJogos, error)
	FindJogoByUsuario(context.Context, *RequestId) (*ListaJogos, error)
	FindJogoFavoritoByUsuario(context.Context, *RequestId) (*ListaJogos, error)
	CreateJogo(context.Context, *Jogo) (*Jogo, error)
	UpdateJogo(context.Context, *Jogo) (*Jogo, error)
	DeleteJogo(context.Context, *RequestId) (*ResponseBool, error)
	mustEmbedUnimplementedJogosServer()
}

// UnimplementedJogosServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJogosServer struct{}

func (UnimplementedJogosServer) FindAllJogos(context.Context, *RequestVazio) (*ListaJogos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllJogos not implemented")
}
func (UnimplementedJogosServer) FindJogoById(context.Context, *RequestId) (*Jogo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJogoById not implemented")
}
func (UnimplementedJogosServer) FindJogoByNome(context.Context, *RequestNome) (*ListaJogos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJogoByNome not implemented")
}
func (UnimplementedJogosServer) FindJogoByGenero(context.Context, *RequestNome) (*ListaJogos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJogoByGenero not implemented")
}
func (UnimplementedJogosServer) FindJogoByUsuario(context.Context, *RequestId) (*ListaJogos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJogoByUsuario not implemented")
}
func (UnimplementedJogosServer) FindJogoFavoritoByUsuario(context.Context, *RequestId) (*ListaJogos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindJogoFavoritoByUsuario not implemented")
}
func (UnimplementedJogosServer) CreateJogo(context.Context, *Jogo) (*Jogo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJogo not implemented")
}
func (UnimplementedJogosServer) UpdateJogo(context.Context, *Jogo) (*Jogo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJogo not implemented")
}
func (UnimplementedJogosServer) DeleteJogo(context.Context, *RequestId) (*ResponseBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJogo not implemented")
}
func (UnimplementedJogosServer) mustEmbedUnimplementedJogosServer() {}
func (UnimplementedJogosServer) testEmbeddedByValue()               {}

// UnsafeJogosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JogosServer will
// result in compilation errors.
type UnsafeJogosServer interface {
	mustEmbedUnimplementedJogosServer()
}

func RegisterJogosServer(s grpc.ServiceRegistrar, srv JogosServer) {
	// If the following call pancis, it indicates UnimplementedJogosServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Jogos_ServiceDesc, srv)
}

func _Jogos_FindAllJogos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVazio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).FindAllJogos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_FindAllJogos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).FindAllJogos(ctx, req.(*RequestVazio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_FindJogoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).FindJogoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_FindJogoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).FindJogoById(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_FindJogoByNome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNome)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).FindJogoByNome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_FindJogoByNome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).FindJogoByNome(ctx, req.(*RequestNome))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_FindJogoByGenero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNome)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).FindJogoByGenero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_FindJogoByGenero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).FindJogoByGenero(ctx, req.(*RequestNome))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_FindJogoByUsuario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).FindJogoByUsuario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_FindJogoByUsuario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).FindJogoByUsuario(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_FindJogoFavoritoByUsuario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).FindJogoFavoritoByUsuario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_FindJogoFavoritoByUsuario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).FindJogoFavoritoByUsuario(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_CreateJogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jogo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).CreateJogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_CreateJogo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).CreateJogo(ctx, req.(*Jogo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_UpdateJogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jogo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).UpdateJogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_UpdateJogo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).UpdateJogo(ctx, req.(*Jogo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jogos_DeleteJogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JogosServer).DeleteJogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jogos_DeleteJogo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JogosServer).DeleteJogo(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// Jogos_ServiceDesc is the grpc.ServiceDesc for Jogos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jogos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Jogos",
	HandlerType: (*JogosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllJogos",
			Handler:    _Jogos_FindAllJogos_Handler,
		},
		{
			MethodName: "FindJogoById",
			Handler:    _Jogos_FindJogoById_Handler,
		},
		{
			MethodName: "FindJogoByNome",
			Handler:    _Jogos_FindJogoByNome_Handler,
		},
		{
			MethodName: "FindJogoByGenero",
			Handler:    _Jogos_FindJogoByGenero_Handler,
		},
		{
			MethodName: "FindJogoByUsuario",
			Handler:    _Jogos_FindJogoByUsuario_Handler,
		},
		{
			MethodName: "FindJogoFavoritoByUsuario",
			Handler:    _Jogos_FindJogoFavoritoByUsuario_Handler,
		},
		{
			MethodName: "CreateJogo",
			Handler:    _Jogos_CreateJogo_Handler,
		},
		{
			MethodName: "UpdateJogo",
			Handler:    _Jogos_UpdateJogo_Handler,
		},
		{
			MethodName: "DeleteJogo",
			Handler:    _Jogos_DeleteJogo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jogos.proto",
}
