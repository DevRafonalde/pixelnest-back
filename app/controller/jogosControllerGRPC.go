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
type JogosServer struct {
	pb.UnimplementedJogosServer
	jogoService         *service.JogoService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewJogosServer(jogoService *service.JogoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *JogosServer {
	return &JogosServer{
		jogoService:         jogoService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (jogoServer *JogosServer) mustEmbedUnimplementedJogosServer() {
}

// Função para buscar por todos os jogos
func (jogoServer *JogosServer) FindAllJogos(context context.Context, req *pb.RequestVazio) (*pb.ListaJogos, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-jogo")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaJogos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numerosTelefonicos, erroService := jogoServer.jogoService.FindAllJogos(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os jogos",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaJogos{Jogos: numerosTelefonicos}, nil
}

// Função para buscar por um jogo pelo ID
func (jogoServer *JogosServer) FindJogoById(context context.Context, req *pb.RequestId) (*pb.Jogo, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-jogo-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Jogo{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	jogo, erroService := jogoServer.jogoService.FindJogoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um jogo pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("jogo", jogo),
	)

	return jogo, nil
}

// Função para buscar por um jogo pelo nome
func (jogoServer *JogosServer) FindJogoByNome(context context.Context, req *pb.RequestNome) (*pb.ListaJogos, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-jogo-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaJogos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	jogos, erroService := jogoServer.jogoService.FindJogoByNome(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados jogos pelo nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaJogos{Jogos: jogos}, nil
}

// Função para buscar por um jogo pelo nome
func (jogoServer *JogosServer) FindJogoByGenero(context context.Context, req *pb.RequestNome) (*pb.ListaJogos, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-jogo-by-genero")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaJogos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	jogos, erroService := jogoServer.jogoService.FindJogoByGenero(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados jogos pelo gênero",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaJogos{Jogos: jogos}, nil
}

// Função para buscar por jogos pelo ID de um usuário
func (jogoServer *JogosServer) FindJogoByUsuario(context context.Context, req *pb.RequestId) (*pb.ListaJogos, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-jogo-by-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaJogos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	jogos, erroService := jogoServer.jogoService.FindJogoByUsuario(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados jogos pelo usuário",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaJogos{Jogos: jogos}, nil
}

// Função para buscar por jogos favoritos de um usuário pelo ID
func (jogoServer *JogosServer) FindJogoFavoritoByUsuario(context context.Context, req *pb.RequestId) (*pb.ListaJogos, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-jogo-by-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaJogos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	jogos, erroService := jogoServer.jogoService.FindJogoFavoritoByUsuario(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados jogos favoritos de um usuário",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaJogos{Jogos: jogos}, nil
}

// Função para criar um novo jogo
func (jogoServer *JogosServer) CreateJogo(context context.Context, req *pb.Jogo) (*pb.Jogo, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-jogo")
	if retornoMiddleware.Erro != nil {
		return &pb.Jogo{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	jogo, erroService := jogoServer.jogoService.CreateJogo(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo jogo",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", jogo),
	)

	return jogo, nil
}

// Função para atualizar um jogo existente no banco de dados
func (jogoServer *JogosServer) UpdateJogo(context context.Context, req *pb.Jogo) (*pb.Jogo, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-jogo")
	if retornoMiddleware.Erro != nil {
		return &pb.Jogo{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	jogoAntigo, erroService := jogoServer.jogoService.FindJogoById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	jogoAtualizado, erroService := jogoServer.jogoService.UpdateJogo(context, req, jogoAntigo)
	if erroService.Erro != nil {
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo jogo via CSV",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", jogoAntigo),
		zap.Any("numeroAtualizado", jogoAtualizado),
	)

	return jogoAtualizado, nil
}

// Função para deletar um jogo existente no banco de dados
func (jogoServer *JogosServer) DeleteJogo(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := jogoServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-jogo-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	jogo, erroService := jogoServer.jogoService.FindJogoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := jogoServer.jogoService.DeleteJogoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		logger.Logger.Error("Não existe jogo com o ID enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe jogo com o ID enviado")
	}

	logger.Logger.Info("Deletado um jogo",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", jogo),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
