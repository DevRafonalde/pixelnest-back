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
	favoritoService     *service.FavoritoService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewFavoritosServer(favoritoService *service.FavoritoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *FavoritosServer {
	return &FavoritosServer{
		favoritoService:     favoritoService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (favoritoServer *FavoritosServer) mustEmbedUnimplementedFavoritosServer() {}

// Função para buscar por um favorito pelo ID
func (favoritoServer *FavoritosServer) FindFavoritoById(context context.Context, req *pb.RequestId) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := favoritoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-favorito-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Favorito{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	favorito, erroService := favoritoServer.favoritoService.FindFavoritoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar favorito pelo ID "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um favorito por ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("favorito", favorito),
	)

	return favorito, nil
}

// Função para buscar por um favorito pelo ID
func (favoritoServer *FavoritosServer) FindJogosFavoritosByUsuario(context context.Context, req *pb.RequestId) (*pb.ListaFavoritos, error) {
	usuarioSolicitante, retornoMiddleware := favoritoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-favorito-by-id-externo")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaFavoritos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	favoritos, erroService := favoritoServer.favoritoService.FindJogosFavoritosByUsuario(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar favorito pelo ID "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um favorito por ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("favorito", favoritos),
	)

	return &pb.ListaFavoritos{Favoritos: favoritos}, nil
}

// Função para buscar por um favorito pelo nome
func (favoritoServer *FavoritosServer) FindProdutosFavoritosByUsuario(context context.Context, req *pb.RequestId) (*pb.ListaFavoritos, error) {
	usuarioSolicitante, retornoMiddleware := favoritoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-favorito-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaFavoritos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	favoritos, erroService := favoritoServer.favoritoService.FindProdutosFavoritosByUsuario(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar o favorito "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um favorito por nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("favorito", favoritos),
	)

	return &pb.ListaFavoritos{Favoritos: favoritos}, nil
}

// Função para criar uma novo favorito
func (favoritoServer *FavoritosServer) CreateFavorito(context context.Context, req *pb.Favorito) (*pb.Favorito, error) {
	usuarioSolicitante, retornoMiddleware := favoritoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-favorito")
	if retornoMiddleware.Erro != nil {
		return &pb.Favorito{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	favoritoCriada, erroService := favoritoServer.favoritoService.CreateFavorito(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao criar o favorito "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo favorito",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("favorito", favoritoCriada),
	)

	return favoritoCriada, nil
}

// Função para deletar um favorito existente no banco
func (favoritoServer *FavoritosServer) DeleteFavorito(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := favoritoServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-favorito-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	favorito, erroService := favoritoServer.favoritoService.FindFavoritoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := favoritoServer.favoritoService.DeleteFavoritoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		logger.Logger.Error("Não existe favorito com o ID enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe favorito com o ID enviado")
	}

	logger.Logger.Info("Deletado um favorito",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("favorito", favorito),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
