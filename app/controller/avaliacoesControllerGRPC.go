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
type AvaliacaosServer struct {
	pb.UnimplementedAvaliacaosServer
	avaliacaoService    *service.AvaliacaoService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewAvaliacaosServer(avaliacaoService *service.AvaliacaoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *AvaliacaosServer {
	return &AvaliacaosServer{
		avaliacaoService:    avaliacaoService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (avaliacoesServer *AvaliacaosServer) mustEmbedUnimplementedAvaliacaosServer() {}

// Função para buscar por todas as avaliacoes
func (avaliacoesServer *AvaliacaosServer) FindAllAvaliacoes(context context.Context, req *pb.RequestVazio) (*pb.ListaAvaliacoes, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-avaliacoes")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaAvaliacoes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacoes, erroService := avaliacoesServer.avaliacaoService.FindAllAvaliacoes(context)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar todas as avaliações "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas todas as avaliações",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaAvaliacoes{Avaliacoes: avaliacoes}, nil
}

// Função para buscar por uma avaliacao pelo ID
func (avaliacoesServer *AvaliacaosServer) FindAvaliacaoById(context context.Context, req *pb.RequestId) (*pb.Avaliacao, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Avaliacao{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	avaliacao, erroService := avaliacoesServer.avaliacaoService.FindAvaliacaoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar avaliacao pelo ID "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma avaliacao por ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacao", avaliacao),
	)

	return avaliacao, nil
}

// Função para buscar por uma avaliacao pelo nome
func (avaliacoesServer *AvaliacaosServer) FindAvaliacaoByUsuario(context context.Context, req *pb.RequestId) (*pb.ListaAvaliacoes, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaAvaliacoes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacao, erroService := avaliacoesServer.avaliacaoService.FindAvaliacaoByUsuario(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar as avaliações pelo usuário "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas as avaliações de um usuário",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacao", avaliacao),
	)

	return &pb.ListaAvaliacoes{Avaliacoes: avaliacao}, nil
}

// Função para buscar por uma avaliacao pelo Produto
func (avaliacoesServer *AvaliacaosServer) FindAvaliacaoByProduto(context context.Context, req *pb.RequestId) (*pb.ListaAvaliacoes, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-produto")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacoes, erroService := avaliacoesServer.avaliacaoService.FindAvaliacaoByProduto(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar as avaliações pelo produto "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas as avaliações pelo Produto",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("produtoId", req.GetID()),
	)

	return &pb.ListaAvaliacoes{Avaliacoes: avaliacoes}, nil
}

// Função para buscar por uma avaliacao pelo Produto
func (avaliacoesServer *AvaliacaosServer) FindAvaliacaoByJogo(context context.Context, req *pb.RequestId) (*pb.ListaAvaliacoes, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-jogo")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacoes, erroService := avaliacoesServer.avaliacaoService.FindAvaliacaoByJogo(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar as avaliações pelo produto "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada a avaliacao pelo código do IBGE",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacao", avaliacoes),
	)

	return &pb.ListaAvaliacoes{Avaliacoes: avaliacoes}, nil
}

// Função para criar uma nova avaliacao
func (avaliacoesServer *AvaliacaosServer) CreateAvaliacao(context context.Context, req *pb.Avaliacao) (*pb.Avaliacao, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-avaliacao")
	if retornoMiddleware.Erro != nil {
		return &pb.Avaliacao{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacaoCriada, erroService := avaliacoesServer.avaliacaoService.CreateAvaliacao(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao criar a avaliacao "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criada uma nova avaliacao",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacao", avaliacaoCriada),
	)

	return avaliacaoCriada, nil
}

// Função para atualizar uma avaliacao já existente no banco
func (avaliacoesServer *AvaliacaosServer) UpdateAvaliacao(context context.Context, avaliacao *pb.Avaliacao) (*pb.Avaliacao, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-avaliacao")
	if retornoMiddleware.Erro != nil {
		return &pb.Avaliacao{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacaoAntiga, erroService := avaliacoesServer.avaliacaoService.FindAvaliacaoById(context, avaliacao.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	avaliacaoAtualizada, erroService := avaliacoesServer.avaliacaoService.UpdateAvaliacao(context, avaliacao, avaliacaoAntiga)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizada uma avaliacao existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacaoAntiga", avaliacaoAntiga),
		zap.Any("avaliacaoAtualizada", avaliacaoAtualizada),
	)

	return avaliacaoAtualizada, nil
}

// Função para deletar uma avaliacao existente no banco
func (avaliacoesServer *AvaliacaosServer) DeleteAvaliacao(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-avaliacao-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	avaliacao, erroService := avaliacoesServer.avaliacaoService.FindAvaliacaoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := avaliacoesServer.avaliacaoService.DeleteAvaliacaoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		logger.Logger.Error("Não existe avaliacao com o ID enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe avaliacao com o ID enviado")
	}

	logger.Logger.Info("Deletada uma avaliacao",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacao", avaliacao),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
