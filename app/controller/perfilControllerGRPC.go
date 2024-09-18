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
type PerfisServer struct {
	pb.UnimplementedPerfisServer
	perfilService       *service.PerfilService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewPerfisServer(perfilService *service.PerfilService, permissaoMiddleware *middlewares.PermissoesMiddleware) *PerfisServer {
	return &PerfisServer{
		perfilService:       perfilService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (perfisServer *PerfisServer) mustEmbedUnimplementedPerfisServer() {}

// Função para buscar por todos os perfis
func (perfisServer *PerfisServer) FindAllPerfis(context context.Context, req *pb.RequestVazio) (*pb.ListaPerfis, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-perfis")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	perfis, erroService := perfisServer.perfilService.FindAllPerfis(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os perfis",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaPerfis{Perfis: perfis}, nil
}

// Função para buscar por um perfil pelo ID
func (perfisServer *PerfisServer) FindPerfilById(context context.Context, req *pb.RequestId) (*pb.PerfilPermissoes, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "get-perfil-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.PerfilPermissoes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	perfilPermissao, erroService := perfisServer.perfilService.FindPerfilById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um perfil pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return perfilPermissao, nil
}

// Função para buscar por todos os usuários vinculados àquele perfil
func (perfisServer *PerfisServer) GetUsuariosVinculados(context context.Context, req *pb.RequestId) (*pb.ResponseGetUsuariosVinculados, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuarios-vinculados")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseGetUsuariosVinculados{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	usuarios, erroService := perfisServer.perfilService.GetUsuariosVinculados(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados os usuários vinculados a um perfil pelo ID do perfil",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Int32("perfilId", id),
	)

	return &pb.ResponseGetUsuariosVinculados{Usuarios: usuarios}, nil
}

// Função para buscar por todas as permissões vinculadas àquele perfil
func (perfisServer *PerfisServer) GetPermissoesVinculadas(context context.Context, req *pb.RequestId) (*pb.ResponseGetPermissoesVinculadas, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "get-permissoes-vinculadas")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseGetPermissoesVinculadas{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	permissoes, erroService := perfisServer.perfilService.GetPermissoesVinculadas(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas as permissões vinculadas a um perfil pelo ID do perfil",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Int32("perfilId", id),
	)

	return &pb.ResponseGetPermissoesVinculadas{Permissoes: permissoes}, nil
}

// Função para criar um novo perfil
func (perfisServer *PerfisServer) CreatePerfil(context context.Context, req *pb.PerfilPermissoes) (*pb.PerfilPermissoes, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-perfil")
	if retornoMiddleware.Erro != nil {
		return &pb.PerfilPermissoes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	perfilCriado, erroService := perfisServer.perfilService.CreatePerfil(context, req) // Cria o perfil
	if erroService.Erro != nil {                                                       // Se houver erro na criação, retorna erro
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.PerfilPermissoes{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um perfil novo",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("perfilCriado", perfilCriado),
	)

	return perfilCriado, nil
}

// Função para clonar um perfil existente
func (perfisServer *PerfisServer) ClonePerfil(context context.Context, req *pb.RequestId) (*pb.PerfilPermissoes, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "post-clone-perfil")
	if retornoMiddleware.Erro != nil {
		return &pb.PerfilPermissoes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	perfilCriado, erroService := perfisServer.perfilService.ClonePerfil(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.PerfilPermissoes{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Clonado um perfil existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("perfilCriado", perfilCriado),
	)

	return perfilCriado, nil
}

// Função para atualizar um perfil existente
func (perfisServer *PerfisServer) UpdatePerfil(context context.Context, req *pb.PerfilPermissoes) (*pb.PerfilPermissoes, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-perfil")
	if retornoMiddleware.Erro != nil {
		return &pb.PerfilPermissoes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	perfilPermissaoAntigo, erroService := perfisServer.perfilService.FindPerfilById(context, req.GetPerfil().GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	perfilAtualizado, erroService := perfisServer.perfilService.UpdatePerfil(context, req, perfilPermissaoAntigo)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.PerfilPermissoes{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizado um perfil existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("perfilAntigo", perfilPermissaoAntigo),
		zap.Any("perfilAtualizado", perfilAtualizado),
	)

	return perfilAtualizado, nil
}

// Função para ativar um perfil existente
func (perfisServer *PerfisServer) AtivarPerfil(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "put-ativar-perfil")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	perfilPermissao, erroService := perfisServer.perfilService.FindPerfilById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = perfisServer.perfilService.AtivarPerfilById(context, perfilPermissao)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Ativado um perfil existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("perfil", perfilPermissao.GetPerfil()),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para desativar um perfil existente
func (perfisServer *PerfisServer) DesativarPerfil(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := perfisServer.permissaoMiddleware.PermissaoMiddleware(context, "put-desativar-perfil")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	perfilPermissao, erroService := perfisServer.perfilService.FindPerfilById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = perfisServer.perfilService.DesativarPerfilById(context, perfilPermissao)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Desativado um perfil existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("perfil", perfilPermissao.GetPerfil()),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
