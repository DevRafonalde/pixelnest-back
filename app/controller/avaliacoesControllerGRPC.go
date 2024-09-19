package controller

import (
	"context"
	"pixelnest/app/configuration/logger"
	"pixelnest/app/controller/middlewares"
	"pixelnest/app/service"
	"strconv"

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
		logger.Logger.Error("Erro ao buscar todas as avaliacoes "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas todas as avaliacao",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaAvaliacoes{Avaliacoes: avaliacoes}, nil
}

// Função para buscar por uma avaliacao pelo ID
func (avaliacoesServer *AvaliacaosServer) FindCidadeById(context context.Context, req *pb.RequestId) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	avaliacao, erroService := avaliacoesServer.avaliacaoService.FindCidadeById(context, id)
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
func (avaliacoesServer *AvaliacaosServer) FindCidadeByNome(context context.Context, req *pb.RequestNome) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacao, erroService := avaliacoesServer.avaliacaoService.FindCidadeByNome(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar a avaliacao "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma avaliacao por nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacao", avaliacao),
	)

	return avaliacao, nil
}

// Função para buscar por uma avaliacao pela UF
func (avaliacoesServer *AvaliacaosServer) FindCidadeByUF(context context.Context, req *pb.RequestUF) (*pb.ListaAvaliacoes, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-uf")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacoes, erroService := avaliacoesServer.avaliacaoService.FindCidadeByUF(context, req.GetUf())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar as avaliacoes pela UF "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas as avaliacoes pela UF",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("UF", req.GetUf()),
	)

	return &pb.ListaAvaliacoes{Avaliacoes: avaliacoes}, nil
}

// Função para buscar por uma avaliacao pela UF
func (avaliacoesServer *AvaliacaosServer) FindCidadeByCodIbge(context context.Context, req *pb.RequestCodIbge) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-avaliacao-by-codIbge")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacao, erroService := avaliacoesServer.avaliacaoService.FindCidadeByCodIbge(context, req.GetCodIbge())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar a avaliacao pelo código do IBGE "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada a avaliacao pelo código do IBGE",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("avaliacao", avaliacao),
	)

	return avaliacao, nil
}

// Função para criar uma nova avaliacao
func (avaliacoesServer *AvaliacaosServer) CreateCidade(context context.Context, req *pb.Cidade) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-avaliacao")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacaoCriada, erroService := avaliacoesServer.avaliacaoService.CreateCidade(context, req)
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

// Função para criar avaliacoes em lote a partir de um arquivo csv
func (avaliacoesServer *AvaliacaosServer) CreateCidadeCSV(context context.Context, req *pb.ListaAvaliacoes) (*pb.RetornoCSV, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-avaliacao-csv")
	if retornoMiddleware.Erro != nil {
		return &pb.RetornoCSV{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	var response pb.RetornoCSV
	for i, avaliacao := range req.GetAvaliacoes() {
		// Converte o índice em string
		iString := strconv.Itoa(i)

		avaliacaoCriada, erroService := avaliacoesServer.avaliacaoService.CreateCidade(context, avaliacao)
		if erroService.Erro != nil {
			logger.Logger.Error("Erro ao criar a avaliacao "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
			response.Status = append(response.Status, &pb.CSVStatus{
				Linha:    iString,
				Mensagem: erroService.Erro.Error(),
				Status:   int32(erroService.Status),
			})
			continue
		}

		logger.Logger.Info("Criada uma nova avaliacao",
			zap.Any("usuario", usuarioSolicitante.Usuario),
			zap.Any("avaliacao", avaliacaoCriada),
		)

		response.Status = append(response.Status, &pb.CSVStatus{
			Linha:    iString,
			Mensagem: "Sucesso",
			Status:   int32(codes.OK),
		})
	}

	return &response, nil
}

// Função para atualizar uma avaliacao já existente no banco
func (avaliacoesServer *AvaliacaosServer) UpdateCidade(context context.Context, avaliacao *pb.Cidade) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-avaliacao")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	avaliacaoAntiga, erroService := avaliacoesServer.avaliacaoService.FindCidadeById(context, avaliacao.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	avaliacaoAtualizada, erroService := avaliacoesServer.avaliacaoService.UpdateCidade(context, avaliacao, avaliacaoAntiga)
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
func (avaliacoesServer *AvaliacaosServer) DeleteCidade(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := avaliacoesServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-avaliacao-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	avaliacao, erroService := avaliacoesServer.avaliacaoService.FindCidadeById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := avaliacoesServer.avaliacaoService.DeleteCidadeById(context, id)
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
