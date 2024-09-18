// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: perfis.proto

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
	Perfis_FindAllPerfis_FullMethodName           = "/grpc.Perfis/FindAllPerfis"
	Perfis_FindPerfilById_FullMethodName          = "/grpc.Perfis/FindPerfilById"
	Perfis_GetUsuariosVinculados_FullMethodName   = "/grpc.Perfis/GetUsuariosVinculados"
	Perfis_GetPermissoesVinculadas_FullMethodName = "/grpc.Perfis/GetPermissoesVinculadas"
	Perfis_CreatePerfil_FullMethodName            = "/grpc.Perfis/CreatePerfil"
	Perfis_ClonePerfil_FullMethodName             = "/grpc.Perfis/ClonePerfil"
	Perfis_UpdatePerfil_FullMethodName            = "/grpc.Perfis/UpdatePerfil"
	Perfis_AtivarPerfil_FullMethodName            = "/grpc.Perfis/AtivarPerfil"
	Perfis_DesativarPerfil_FullMethodName         = "/grpc.Perfis/DesativarPerfil"
)

// PerfisClient is the client API for Perfis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Definição do serviço
type PerfisClient interface {
	FindAllPerfis(ctx context.Context, in *RequestVazio, opts ...grpc.CallOption) (*ListaPerfis, error)
	FindPerfilById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*PerfilPermissoes, error)
	GetUsuariosVinculados(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseGetUsuariosVinculados, error)
	GetPermissoesVinculadas(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseGetPermissoesVinculadas, error)
	CreatePerfil(ctx context.Context, in *PerfilPermissoes, opts ...grpc.CallOption) (*PerfilPermissoes, error)
	ClonePerfil(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*PerfilPermissoes, error)
	UpdatePerfil(ctx context.Context, in *PerfilPermissoes, opts ...grpc.CallOption) (*PerfilPermissoes, error)
	AtivarPerfil(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error)
	DesativarPerfil(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error)
}

type perfisClient struct {
	cc grpc.ClientConnInterface
}

func NewPerfisClient(cc grpc.ClientConnInterface) PerfisClient {
	return &perfisClient{cc}
}

func (c *perfisClient) FindAllPerfis(ctx context.Context, in *RequestVazio, opts ...grpc.CallOption) (*ListaPerfis, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListaPerfis)
	err := c.cc.Invoke(ctx, Perfis_FindAllPerfis_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) FindPerfilById(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*PerfilPermissoes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PerfilPermissoes)
	err := c.cc.Invoke(ctx, Perfis_FindPerfilById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) GetUsuariosVinculados(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseGetUsuariosVinculados, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseGetUsuariosVinculados)
	err := c.cc.Invoke(ctx, Perfis_GetUsuariosVinculados_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) GetPermissoesVinculadas(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseGetPermissoesVinculadas, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseGetPermissoesVinculadas)
	err := c.cc.Invoke(ctx, Perfis_GetPermissoesVinculadas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) CreatePerfil(ctx context.Context, in *PerfilPermissoes, opts ...grpc.CallOption) (*PerfilPermissoes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PerfilPermissoes)
	err := c.cc.Invoke(ctx, Perfis_CreatePerfil_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) ClonePerfil(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*PerfilPermissoes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PerfilPermissoes)
	err := c.cc.Invoke(ctx, Perfis_ClonePerfil_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) UpdatePerfil(ctx context.Context, in *PerfilPermissoes, opts ...grpc.CallOption) (*PerfilPermissoes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PerfilPermissoes)
	err := c.cc.Invoke(ctx, Perfis_UpdatePerfil_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) AtivarPerfil(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseBool)
	err := c.cc.Invoke(ctx, Perfis_AtivarPerfil_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perfisClient) DesativarPerfil(ctx context.Context, in *RequestId, opts ...grpc.CallOption) (*ResponseBool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseBool)
	err := c.cc.Invoke(ctx, Perfis_DesativarPerfil_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PerfisServer is the server API for Perfis service.
// All implementations must embed UnimplementedPerfisServer
// for forward compatibility.
//
// Definição do serviço
type PerfisServer interface {
	FindAllPerfis(context.Context, *RequestVazio) (*ListaPerfis, error)
	FindPerfilById(context.Context, *RequestId) (*PerfilPermissoes, error)
	GetUsuariosVinculados(context.Context, *RequestId) (*ResponseGetUsuariosVinculados, error)
	GetPermissoesVinculadas(context.Context, *RequestId) (*ResponseGetPermissoesVinculadas, error)
	CreatePerfil(context.Context, *PerfilPermissoes) (*PerfilPermissoes, error)
	ClonePerfil(context.Context, *RequestId) (*PerfilPermissoes, error)
	UpdatePerfil(context.Context, *PerfilPermissoes) (*PerfilPermissoes, error)
	AtivarPerfil(context.Context, *RequestId) (*ResponseBool, error)
	DesativarPerfil(context.Context, *RequestId) (*ResponseBool, error)
	mustEmbedUnimplementedPerfisServer()
}

// UnimplementedPerfisServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPerfisServer struct{}

func (UnimplementedPerfisServer) FindAllPerfis(context.Context, *RequestVazio) (*ListaPerfis, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllPerfis not implemented")
}
func (UnimplementedPerfisServer) FindPerfilById(context.Context, *RequestId) (*PerfilPermissoes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPerfilById not implemented")
}
func (UnimplementedPerfisServer) GetUsuariosVinculados(context.Context, *RequestId) (*ResponseGetUsuariosVinculados, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsuariosVinculados not implemented")
}
func (UnimplementedPerfisServer) GetPermissoesVinculadas(context.Context, *RequestId) (*ResponseGetPermissoesVinculadas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissoesVinculadas not implemented")
}
func (UnimplementedPerfisServer) CreatePerfil(context.Context, *PerfilPermissoes) (*PerfilPermissoes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePerfil not implemented")
}
func (UnimplementedPerfisServer) ClonePerfil(context.Context, *RequestId) (*PerfilPermissoes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClonePerfil not implemented")
}
func (UnimplementedPerfisServer) UpdatePerfil(context.Context, *PerfilPermissoes) (*PerfilPermissoes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerfil not implemented")
}
func (UnimplementedPerfisServer) AtivarPerfil(context.Context, *RequestId) (*ResponseBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtivarPerfil not implemented")
}
func (UnimplementedPerfisServer) DesativarPerfil(context.Context, *RequestId) (*ResponseBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DesativarPerfil not implemented")
}
func (UnimplementedPerfisServer) mustEmbedUnimplementedPerfisServer() {}
func (UnimplementedPerfisServer) testEmbeddedByValue()                {}

// UnsafePerfisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PerfisServer will
// result in compilation errors.
type UnsafePerfisServer interface {
	mustEmbedUnimplementedPerfisServer()
}

func RegisterPerfisServer(s grpc.ServiceRegistrar, srv PerfisServer) {
	// If the following call pancis, it indicates UnimplementedPerfisServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Perfis_ServiceDesc, srv)
}

func _Perfis_FindAllPerfis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVazio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).FindAllPerfis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_FindAllPerfis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).FindAllPerfis(ctx, req.(*RequestVazio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_FindPerfilById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).FindPerfilById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_FindPerfilById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).FindPerfilById(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_GetUsuariosVinculados_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).GetUsuariosVinculados(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_GetUsuariosVinculados_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).GetUsuariosVinculados(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_GetPermissoesVinculadas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).GetPermissoesVinculadas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_GetPermissoesVinculadas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).GetPermissoesVinculadas(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_CreatePerfil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerfilPermissoes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).CreatePerfil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_CreatePerfil_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).CreatePerfil(ctx, req.(*PerfilPermissoes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_ClonePerfil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).ClonePerfil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_ClonePerfil_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).ClonePerfil(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_UpdatePerfil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerfilPermissoes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).UpdatePerfil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_UpdatePerfil_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).UpdatePerfil(ctx, req.(*PerfilPermissoes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_AtivarPerfil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).AtivarPerfil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_AtivarPerfil_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).AtivarPerfil(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Perfis_DesativarPerfil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfisServer).DesativarPerfil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Perfis_DesativarPerfil_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfisServer).DesativarPerfil(ctx, req.(*RequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// Perfis_ServiceDesc is the grpc.ServiceDesc for Perfis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Perfis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Perfis",
	HandlerType: (*PerfisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllPerfis",
			Handler:    _Perfis_FindAllPerfis_Handler,
		},
		{
			MethodName: "FindPerfilById",
			Handler:    _Perfis_FindPerfilById_Handler,
		},
		{
			MethodName: "GetUsuariosVinculados",
			Handler:    _Perfis_GetUsuariosVinculados_Handler,
		},
		{
			MethodName: "GetPermissoesVinculadas",
			Handler:    _Perfis_GetPermissoesVinculadas_Handler,
		},
		{
			MethodName: "CreatePerfil",
			Handler:    _Perfis_CreatePerfil_Handler,
		},
		{
			MethodName: "ClonePerfil",
			Handler:    _Perfis_ClonePerfil_Handler,
		},
		{
			MethodName: "UpdatePerfil",
			Handler:    _Perfis_UpdatePerfil_Handler,
		},
		{
			MethodName: "AtivarPerfil",
			Handler:    _Perfis_AtivarPerfil_Handler,
		},
		{
			MethodName: "DesativarPerfil",
			Handler:    _Perfis_DesativarPerfil_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "perfis.proto",
}
