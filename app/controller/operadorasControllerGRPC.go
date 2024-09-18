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
type OperadorasServer struct {
	pb.UnimplementedOperadorasServer
	operadoraService    *service.OperadoraService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewOperadorasServer(operadoraService *service.OperadoraService, permissaoMiddleware *middlewares.PermissoesMiddleware) *OperadorasServer {
	return &OperadorasServer{
		operadoraService:    operadoraService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (operadoraServer *OperadorasServer) mustEmbedUnimplementedOperadorasServer() {}

// Função para buscar por todas as operadoras existentes no banco de dados
func (operadoraServer *OperadorasServer) FindAllOperadoras(context context.Context, req *pb.RequestVazio) (*pb.ListaOperadoras, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-operadoras")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaOperadoras{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	operadoras, erroService := operadoraServer.operadoraService.FindAllOperadoras(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscadas todas as operadoras",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaOperadoras{Operadoras: operadoras}, nil
}

// Função para buscar por uma operadora pelo ID
func (operadoraServer *OperadorasServer) FindOperadoraById(context context.Context, req *pb.RequestId) (*pb.Operadora, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "get-operadora-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Operadora{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	operadora, erroService := operadoraServer.operadoraService.FindOperadoraById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma operadora pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("operadora", operadora),
	)

	return operadora, nil
}

// Função para buscar por uma operadora pelo nome
func (operadoraServer *OperadorasServer) FindOperadoraByNome(context context.Context, req *pb.RequestNome) (*pb.Operadora, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "get-operadora-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.Operadora{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	operadora, erroService := operadoraServer.operadoraService.FindOperadoraByNome(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma operadora pelo nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("operadora", operadora),
	)

	return operadora, nil
}

// Função para buscar por uma operadora pela abreviação
func (operadoraServer *OperadorasServer) FindOperadoraByAbreviacao(context context.Context, req *pb.RequestAbreviacao) (*pb.Operadora, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "get-operadora-by-abreviacao")
	if retornoMiddleware.Erro != nil {
		return &pb.Operadora{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	operadora, erroService := operadoraServer.operadoraService.FindOperadoraByAbreviacao(context, req.GetAbreviacao())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma operadora pela abreviação",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("operadora", operadora),
	)

	return operadora, nil
}

// Função para criar uma nova operadora
func (operadoraServer *OperadorasServer) CreateOperadora(context context.Context, req *pb.Operadora) (*pb.Operadora, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-operadora")
	if retornoMiddleware.Erro != nil {
		return &pb.Operadora{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	operadoraCriada, erroService := operadoraServer.operadoraService.CreateOperadora(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criada uma nova operadora",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("operadora", operadoraCriada),
	)

	return operadoraCriada, nil
}

// Função para criar números telefônicos em lote a partir de um arquivo csv
func (operadoraServer *OperadorasServer) CreateOperadoraCSV(context context.Context, req *pb.ListaOperadoras) (*pb.RetornoCSV, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-operadora-csv")
	if retornoMiddleware.Erro != nil {
		return &pb.RetornoCSV{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	var response pb.RetornoCSV
	for i, numeroTelefonico := range req.GetOperadoras() {
		// Converte o índice em string
		iString := strconv.Itoa(i)

		operadoraCriada, erroService := operadoraServer.operadoraService.CreateOperadora(context, numeroTelefonico)
		if erroService.Erro != nil {
			logger.Logger.Error("Erro ao criar a operadora "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
			response.Status = append(response.Status, &pb.CSVStatus{
				Linha:    iString,
				Mensagem: erroService.Erro.Error(),
				Status:   int32(erroService.Status),
			})
			continue
		}

		logger.Logger.Info("Criado uma nova operadora",
			zap.Any("usuario", usuarioSolicitante.Usuario),
			zap.Any("operadora", operadoraCriada),
		)

		response.Status = append(response.Status, &pb.CSVStatus{
			Linha:    iString,
			Mensagem: "Sucesso",
			Status:   int32(codes.OK),
		})
	}

	return &response, nil
}

// Função para atualizar uma operadora já existente no banco
func (operadoraServer *OperadorasServer) UpdateOperadora(context context.Context, req *pb.Operadora) (*pb.Operadora, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-operadora")
	if retornoMiddleware.Erro != nil {
		return &pb.Operadora{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	operadoraAntiga, erroService := operadoraServer.operadoraService.FindOperadoraById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	operadoraCriada, erroService := operadoraServer.operadoraService.UpdateOperadora(context, req, operadoraAntiga)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criada uma nova operadora via CSV",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("operadora", operadoraAntiga),
		zap.Any("operadoraAtualizada", operadoraCriada),
	)

	return operadoraCriada, nil
}

// Função para deletar uma operadora existente no banco
func (operadoraServer *OperadorasServer) DeleteOperadora(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := operadoraServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-operadora-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	operadora, erroService := operadoraServer.operadoraService.FindOperadoraById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := operadoraServer.operadoraService.DeleteOperadoraById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe operadora com o ID enviado")
	}

	logger.Logger.Info("Deletada uma operadora",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("operadora", operadora),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
