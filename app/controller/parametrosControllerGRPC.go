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
type ParametrosServer struct {
	pb.UnimplementedParametrosServer
	parametroService    *service.ParametroService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewParametrosServer(parametroService *service.ParametroService, permissaoMiddleware *middlewares.PermissoesMiddleware) *ParametrosServer {
	return &ParametrosServer{
		parametroService:    parametroService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (parametroServer *ParametrosServer) mustEmbedUnimplementedParametrosServer() {}

// Função para buscar por todas as parametros existentes no banco de dados
func (parametroServer *ParametrosServer) FindAllParametros(context context.Context, req *pb.RequestVazio) (*pb.ListaParametros, error) {
	usuarioSolicitante, retornoMiddleware := parametroServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-parametros")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaParametros{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	parametros, erroService := parametroServer.parametroService.FindAllParametros(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os parâmetros",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaParametros{Parametros: parametros}, nil
}

// Função para buscar por uma parametro pelo nome
func (parametroServer *ParametrosServer) FindParametroByNome(context context.Context, req *pb.RequestNome) (*pb.Parametro, error) {
	usuarioSolicitante, retornoMiddleware := parametroServer.permissaoMiddleware.PermissaoMiddleware(context, "get-parametro-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.Parametro{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	parametro, erroService := parametroServer.parametroService.FindParametroByNome(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um parâmetro pelo nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("parametro", parametro),
	)

	return parametro, nil
}

// Função para buscar por uma parametro pelo nome
func (parametroServer *ParametrosServer) FindParametroById(context context.Context, req *pb.RequestId) (*pb.Parametro, error) {
	usuarioSolicitante, retornoMiddleware := parametroServer.permissaoMiddleware.PermissaoMiddleware(context, "get-parametro-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.Parametro{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	parametro, erroService := parametroServer.parametroService.FindParametroById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um parâmetro pelo id",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("parametro", parametro),
	)

	return parametro, nil
}

// Função para criar uma nova parametro
func (parametroServer *ParametrosServer) CreateParametro(context context.Context, req *pb.Parametro) (*pb.Parametro, error) {
	usuarioSolicitante, retornoMiddleware := parametroServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-parametro")
	if retornoMiddleware.Erro != nil {
		return &pb.Parametro{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	parametroCriado, erroService := parametroServer.parametroService.CreateParametro(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo parâmetro",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("parametro", parametroCriado),
	)

	return parametroCriado, nil
}

// Função para atualizar uma parametro já existente no banco
func (parametroServer *ParametrosServer) UpdateParametro(context context.Context, req *pb.Parametro) (*pb.Parametro, error) {
	usuarioSolicitante, retornoMiddleware := parametroServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-parametro")
	if retornoMiddleware.Erro != nil {
		return &pb.Parametro{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	parametroAntigo, erroService := parametroServer.parametroService.FindParametroById(context, req.GetId())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	parametroAtualizado, erroService := parametroServer.parametroService.UpdateParametro(context, req, parametroAntigo)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizado um parâmetro existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("parametroAntigo", parametroAntigo),
		zap.Any("parametroAtualizado", parametroAtualizado),
	)

	return parametroAtualizado, nil
}

// Função para deletar uma parametro existente no banco
func (parametroServer *ParametrosServer) DeleteParametro(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := parametroServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-parametro-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	parametro, erroService := parametroServer.parametroService.FindParametroById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := parametroServer.parametroService.DeleteParametroById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe parâmetro com o ID enviado")
	}

	logger.Logger.Info("Deletado um parametro",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("parametro", parametro),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
