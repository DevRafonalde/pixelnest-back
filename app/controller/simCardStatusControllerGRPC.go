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
type SimCardStatusServer struct {
	pb.UnimplementedSimCardsStatusServer
	simCardStatusService *service.SimCardStatusService
	permissaoMiddleware  *middlewares.PermissoesMiddleware
}

func NewSimCardStatusServer(simCardStatusService *service.SimCardStatusService, permissaoMiddleware *middlewares.PermissoesMiddleware) *SimCardStatusServer {
	return &SimCardStatusServer{
		simCardStatusService: simCardStatusService,
		permissaoMiddleware:  permissaoMiddleware,
	}
}

func (simCardStatusServer *SimCardStatusServer) mustEmbedUnimplementedSimCardStatusServer() {}

// Função para buscar por todos os statuss de SimCard
func (simCardStatusServer *SimCardStatusServer) FindAllSimCardStatus(context context.Context, req *pb.RequestVazio) (*pb.ListaSimCardStatus, error) {
	usuarioSolicitante, retornoMiddleware := simCardStatusServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-simcardstatus")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaSimCardStatus{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCardStatus, erroService := simCardStatusServer.simCardStatusService.FindAllSimCardStatus(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os status de SimCards",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaSimCardStatus{SimCardStatus: simCardStatus}, nil
}

// Função para buscar por um status de SimCard pelo ID
func (simCardStatusServer *SimCardStatusServer) FindSimCardStatusById(context context.Context, req *pb.RequestId) (*pb.SimCardStatus, error) {
	usuarioSolicitante, retornoMiddleware := simCardStatusServer.permissaoMiddleware.PermissaoMiddleware(context, "get-simcardstatus-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCardStatus{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	simCardStatus, erroService := simCardStatusServer.simCardStatusService.FindSimCardStatusById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um status de SimCards pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("status", simCardStatus),
	)

	return simCardStatus, nil
}

// Função para buscar um status de SimCard pelo status
func (simCardStatusServer *SimCardStatusServer) FindSimCardStatusByNome(context context.Context, req *pb.RequestNome) (*pb.SimCardStatus, error) {
	usuarioSolicitante, retornoMiddleware := simCardStatusServer.permissaoMiddleware.PermissaoMiddleware(context, "get-simcardstatus-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCardStatus{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCardStatus, erroService := simCardStatusServer.simCardStatusService.FindSimCardStatusByStatus(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um status de SimCards pelo nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("status", simCardStatus),
	)

	return simCardStatus, nil
}

// Função para criar um novo status de SimCard
func (simCardStatusServer *SimCardStatusServer) CreateSimCardStatus(context context.Context, req *pb.SimCardStatus) (*pb.SimCardStatus, error) {
	usuarioSolicitante, retornoMiddleware := simCardStatusServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-simcardstatus")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCardStatus{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCardStatus, erroService := simCardStatusServer.simCardStatusService.CreateSimCardStatus(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um status de SimCards pelo nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("status", simCardStatus),
	)

	return simCardStatus, nil
}

// Função para criar novos statuss de SimCard em lote a partir de um arquivo CSV
func (simCardStatusServer *SimCardStatusServer) CreateSimCardStatusCSV(context context.Context, req *pb.ListaSimCardStatus) (*pb.RetornoCSV, error) {
	usuarioSolicitante, retornoMiddleware := simCardStatusServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-simcardstatus-csv")
	if retornoMiddleware.Erro != nil {
		return &pb.RetornoCSV{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	var response pb.RetornoCSV
	for i, simCardStatus := range req.GetSimCardStatus() {
		// Converte o índice em string
		iString := strconv.Itoa(i)

		simCardCriado, erroService := simCardStatusServer.simCardStatusService.CreateSimCardStatus(context, simCardStatus)
		if erroService.Erro != nil {
			logger.Logger.Error("Erro ao criar o status de SimCard "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
			response.Status = append(response.Status, &pb.CSVStatus{
				Linha:    iString,
				Mensagem: erroService.Erro.Error(),
				Status:   int32(erroService.Status),
			})
			continue
		}

		logger.Logger.Info("Criado um novo status de SimCard",
			zap.Any("usuario", usuarioSolicitante.Usuario),
			zap.Any("simCardStatus", simCardCriado),
		)

		response.Status = append(response.Status, &pb.CSVStatus{
			Linha:    iString,
			Mensagem: "Sucesso",
			Status:   int32(codes.OK),
		})
	}

	return &response, nil
}

// Função para atualizar um status de SimCard já existente
func (simCardStatusServer *SimCardStatusServer) UpdateSimCardStatus(context context.Context, req *pb.SimCardStatus) (*pb.SimCardStatus, error) {
	usuarioSolicitante, retornoMiddleware := simCardStatusServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-simcardstatus")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCardStatus{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCardStatusAntigo, erroService := simCardStatusServer.simCardStatusService.FindSimCardStatusById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	simCardStatus, erroService := simCardStatusServer.simCardStatusService.UpdateSimCardStatus(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizado um SimCard existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simCardStatusAntigo", simCardStatusAntigo),
		zap.Any("simCardStatusAtualizado", simCardStatus),
	)

	return simCardStatus, nil
}

// Função para deletar um status de SimCard existente
func (simCardStatusServer *SimCardStatusServer) DeleteSimCardStatus(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := simCardStatusServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-simcardstatus-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	simCardStatus, erroService := simCardStatusServer.simCardStatusService.FindSimCardStatusById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := simCardStatusServer.simCardStatusService.DeleteSimCardStatusById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.Aborted, "Não existe SimCard com o ID enviado")
	}

	logger.Logger.Info("Deletado um status SimCard",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", simCardStatus),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
