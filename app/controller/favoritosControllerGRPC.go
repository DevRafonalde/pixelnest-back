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
type FavoritosServer struct {
	pb.UnimplementedFavoritosServer
	clienteService      *service.FavoritoService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewFavoritosServer(clienteService *service.FavoritoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *FavoritosServer {
	return &FavoritosServer{
		clienteService:      clienteService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (clienteServer *FavoritosServer) mustEmbedUnimplementedFavoritosServer() {}

// Função para buscar por todos os clientes
func (clienteServer *FavoritosServer) FindAllFavoritos(context context.Context, req *pb.RequestVazio) (*pb.ListaFavoritos, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-clientes")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaFavoritos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	clientes, erroService := clienteServer.clienteService.FindAllFavoritos(context)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar todos os clientes "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os clientes",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaFavoritos{Favoritos: clientes}, nil
}

// Função para buscar por um cliente pelo ID
func (clienteServer *FavoritosServer) FindFavoritoById(context context.Context, req *pb.RequestId) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Favorito{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cliente, erroService := clienteServer.clienteService.FindFavoritoById(context, id)
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
func (clienteServer *FavoritosServer) FindFavoritoByIdExterno(context context.Context, req *pb.RequestId) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-id-externo")
	if retornoMiddleware.Erro != nil {
		return &pb.Favorito{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cliente, erroService := clienteServer.clienteService.FindFavoritoByIdExterno(context, id)
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
func (clienteServer *FavoritosServer) FindFavoritoByNome(context context.Context, req *pb.RequestNome) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.Favorito{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cliente, erroService := clienteServer.clienteService.FindFavoritoByNome(context, req.GetNome())
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
func (clienteServer *FavoritosServer) FindFavoritoByDocumento(context context.Context, req *pb.RequestDocumento) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-documento")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cliente, erroService := clienteServer.clienteService.FindFavoritoByDocumento(context, req.GetDocumento())
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
func (clienteServer *FavoritosServer) FindFavoritoByCodIbge(context context.Context, req *pb.RequestCodReserva) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cliente-by-codReserva")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cliente, erroService := clienteServer.clienteService.FindFavoritoByCodReserva(context, req.GetCodReserva())
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
func (clienteServer *FavoritosServer) CreateFavorito(context context.Context, req *pb.Favorito) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-cliente")
	if retornoMiddleware.Erro != nil {
		return &pb.Favorito{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	clienteCriada, erroService := clienteServer.clienteService.CreateFavorito(context, req)
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
func (clienteServer *FavoritosServer) UpdateFavorito(context context.Context, cliente *pb.Favorito) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-cliente")
	if retornoMiddleware.Erro != nil {
		return &pb.Favorito{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	clienteAntiga, erroService := clienteServer.clienteService.FindFavoritoById(context, cliente.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	clienteAtualizado, erroService := clienteServer.clienteService.UpdateFavorito(context, cliente, clienteAntiga)
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
func (clienteServer *FavoritosServer) DeleteFavorito(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := clienteServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-cliente-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cliente, erroService := clienteServer.clienteService.FindFavoritoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := clienteServer.clienteService.DeleteFavoritoById(context, id)
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
