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
type CidadesServer struct {
	pb.UnimplementedCidadesServer
	cidadeService       *service.CidadeService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewCidadesServer(cidadeService *service.CidadeService, permissaoMiddleware *middlewares.PermissoesMiddleware) *CidadesServer {
	return &CidadesServer{
		cidadeService:       cidadeService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (cidadeServer *CidadesServer) mustEmbedUnimplementedCidadesServer() {}

// Função para buscar por todas as cidades
func (cidadeServer *CidadesServer) FindAllCidades(context context.Context, req *pb.RequestVazio) (*pb.ListaCidades, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-cidades")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaCidades{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cidades, erroService := cidadeServer.cidadeService.FindAllCidades(context)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar todas as cidades "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas todas as cidade",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaCidades{Cidades: cidades}, nil
}

// Função para buscar por uma cidade pelo ID
func (cidadeServer *CidadesServer) FindCidadeById(context context.Context, req *pb.RequestId) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cidade-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cidade, erroService := cidadeServer.cidadeService.FindCidadeById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar cidade pelo ID "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma cidade por ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cidade", cidade),
	)

	return cidade, nil
}

// Função para buscar por uma cidade pelo nome
func (cidadeServer *CidadesServer) FindCidadeByNome(context context.Context, req *pb.RequestNome) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cidade-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cidade, erroService := cidadeServer.cidadeService.FindCidadeByNome(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar a cidade "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma cidade por nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cidade", cidade),
	)

	return cidade, nil
}

// Função para buscar por uma cidade pela UF
func (cidadeServer *CidadesServer) FindCidadeByUF(context context.Context, req *pb.RequestUF) (*pb.ListaCidades, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cidade-by-uf")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cidades, erroService := cidadeServer.cidadeService.FindCidadeByUF(context, req.GetUf())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar as cidades pela UF "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas as cidades pela UF",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("UF", req.GetUf()),
	)

	return &pb.ListaCidades{Cidades: cidades}, nil
}

// Função para buscar por uma cidade pela UF
func (cidadeServer *CidadesServer) FindCidadeByCodIbge(context context.Context, req *pb.RequestCodIbge) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "get-cidade-by-codIbge")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cidade, erroService := cidadeServer.cidadeService.FindCidadeByCodIbge(context, req.GetCodIbge())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar a cidade pelo código do IBGE "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada a cidade pelo código do IBGE",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cidade", cidade),
	)

	return cidade, nil
}

// Função para criar uma nova cidade
func (cidadeServer *CidadesServer) CreateCidade(context context.Context, req *pb.Cidade) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-cidade")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cidadeCriada, erroService := cidadeServer.cidadeService.CreateCidade(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao criar a cidade "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criada uma nova cidade",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cidade", cidadeCriada),
	)

	return cidadeCriada, nil
}

// Função para criar cidades em lote a partir de um arquivo csv
func (cidadeServer *CidadesServer) CreateCidadeCSV(context context.Context, req *pb.ListaCidades) (*pb.RetornoCSV, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-cidade-csv")
	if retornoMiddleware.Erro != nil {
		return &pb.RetornoCSV{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	var response pb.RetornoCSV
	for i, cidade := range req.GetCidades() {
		// Converte o índice em string
		iString := strconv.Itoa(i)

		cidadeCriada, erroService := cidadeServer.cidadeService.CreateCidade(context, cidade)
		if erroService.Erro != nil {
			logger.Logger.Error("Erro ao criar a cidade "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
			response.Status = append(response.Status, &pb.CSVStatus{
				Linha:    iString,
				Mensagem: erroService.Erro.Error(),
				Status:   int32(erroService.Status),
			})
			continue
		}

		logger.Logger.Info("Criada uma nova cidade",
			zap.Any("usuario", usuarioSolicitante.Usuario),
			zap.Any("cidade", cidadeCriada),
		)

		response.Status = append(response.Status, &pb.CSVStatus{
			Linha:    iString,
			Mensagem: "Sucesso",
			Status:   int32(codes.OK),
		})
	}

	return &response, nil
}

// Função para atualizar uma cidade já existente no banco
func (cidadeServer *CidadesServer) UpdateCidade(context context.Context, cidade *pb.Cidade) (*pb.Cidade, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-cidade")
	if retornoMiddleware.Erro != nil {
		return &pb.Cidade{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	cidadeAntiga, erroService := cidadeServer.cidadeService.FindCidadeById(context, cidade.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	cidadeAtualizada, erroService := cidadeServer.cidadeService.UpdateCidade(context, cidade, cidadeAntiga)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizada uma cidade existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cidadeAntiga", cidadeAntiga),
		zap.Any("cidadeAtualizada", cidadeAtualizada),
	)

	return cidadeAtualizada, nil
}

// Função para deletar uma cidade existente no banco
func (cidadeServer *CidadesServer) DeleteCidade(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := cidadeServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-cidade-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	cidade, erroService := cidadeServer.cidadeService.FindCidadeById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := cidadeServer.cidadeService.DeleteCidadeById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		logger.Logger.Error("Não existe cidade com o ID enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe cidade com o ID enviado")
	}

	logger.Logger.Info("Deletada uma cidade",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("cidade", cidade),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
