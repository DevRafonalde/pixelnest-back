package controller

import (
	"context"
	"pixelnest/app/configuration/logger"
	"pixelnest/app/controller/middlewares"
	"pixelnest/app/service"

	pb "pixelnest/app/model/grpc" // Importa o pacote gerado pelos arquivos .proto

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implementação do servidor
type ClientesServer struct {
	pb.UnimplementedClientesServer
	clienteService      *service.ClienteService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewClientesServer(clienteService *service.ClienteService, permissaoMiddleware *middlewares.PermissoesMiddleware) *ClientesServer {
	return &ClientesServer{
		clienteService:      clienteService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (clienteServer *ClientesServer) mustEmbedUnimplementedClientesServer() {}

// Função para buscar por todos os clientes
func (clienteServer *ClientesServer) FindAllClientes(context context.Context, req *pb.RequestVazio) (*pb.ListaClientes, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-clientes")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaClientes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	clientes, erroService := clienteServer.clienteService.FindAllClientes(context)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar todos os clientes "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os clientes",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaClientes{Clientes: clientes}, nil
}

// Função para buscar por um cliente pelo ID
func (clienteServer *ClientesServer) FindClienteById(context context.Context, req *pb.RequestId) (*pb.Cliente, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Cliente{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cliente, erroService := clienteServer.clienteService.FindClienteById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar cliente pelo ID "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um cliente por ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cliente", cliente),
	)

	return cliente, nil
}

// Função para buscar por um cliente pelo ID
func (clienteServer *ClientesServer) FindClienteByIdExterno(context context.Context, req *pb.RequestId) (*pb.Cliente, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-id-externo")
	if retornoMiddleware.Erro != nil {
		return &pb.Cliente{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cliente, erroService := clienteServer.clienteService.FindClienteByIdExterno(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar cliente pelo ID "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um cliente por ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cliente", cliente),
	)

	return cliente, nil
}

// Função para buscar por um cliente pelo nome
func (clienteServer *ClientesServer) FindClienteByNome(context context.Context, req *pb.RequestNome) (*pb.Cliente, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.Cliente{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cliente, erroService := clienteServer.clienteService.FindClienteByNome(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar o cliente "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um cliente por nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cliente", cliente),
	)

	return cliente, nil
}

// Função para buscar por um cliente pela UF
func (clienteServer *ClientesServer) FindClienteByDocumento(context context.Context, req *pb.RequestDocumento) (*pb.Cliente, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-documento")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cliente, erroService := clienteServer.clienteService.FindClienteByDocumento(context, req.GetDocumento())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar os clientes pela UF "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados os clientes pela UF",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("documento", req.GetDocumento()),
	)

	return cliente, nil
}

// Função para buscar por um cliente pela UF
func (clienteServer *ClientesServer) FindClienteByCodIbge(context context.Context, req *pb.RequestCodReserva) (*pb.Cliente, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-codReserva")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cliente, erroService := clienteServer.clienteService.FindClienteByCodReserva(context, req.GetCodReserva())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar o cliente pelo código de reserva "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado o cliente pelo código de reserva",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cliente", cliente),
	)

	return cliente, nil
}

// Função para criar uma novo cliente
func (clienteServer *ClientesServer) CreateCliente(context context.Context, req *pb.Cliente) (*pb.Cliente, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-cliente")
	if retornoMiddleware.Erro != nil {
		return &pb.Cliente{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	clienteCriada, erroService := clienteServer.clienteService.CreateCliente(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao criar o cliente "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo cliente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cliente", clienteCriada),
	)

	return clienteCriada, nil
}

// Função para atualizar um cliente já existente no banco
func (clienteServer *ClientesServer) UpdateCliente(context context.Context, cliente *pb.Cliente) (*pb.Cliente, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-cliente")
	if retornoMiddleware.Erro != nil {
		return &pb.Cliente{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	clienteAntiga, erroService := clienteServer.clienteService.FindClienteById(context, cliente.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	clienteAtualizado, erroService := clienteServer.clienteService.UpdateCliente(context, cliente, clienteAntiga)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizado um cliente existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("clienteAntiga", clienteAntiga),
		zap.Any("clienteAtualizado", clienteAtualizado),
	)

	return clienteAtualizado, nil
}

// Função para deletar um cliente existente no banco
func (clienteServer *ClientesServer) DeleteCliente(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-cliente-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cliente, erroService := clienteServer.clienteService.FindClienteById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := clienteServer.clienteService.DeleteClienteById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		logger.Logger.Error("Não existe cliente com o ID enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe cliente com o ID enviado")
	}

	logger.Logger.Info("Deletado um cliente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cliente", cliente),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
