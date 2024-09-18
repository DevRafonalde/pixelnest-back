package controller

import (
	"context"
	"os"
	"pixelnest/app/configuration/logger"
	"pixelnest/app/controller/middlewares"
	"pixelnest/app/service"
	"strconv"
	"strings"
	"time"

	pb "pixelnest/app/model/grpc" // Importa o pacote gerado pelos arquivos .proto

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implementação do servidor
type SimCardsServer struct {
	pb.UnimplementedSimCardsServer
	simCardService          *service.SimCardService
	numeroTelefonicoService *service.NumeroTelefonicoService
	permissaoMiddleware     *middlewares.PermissoesMiddleware
}

func NewSimCardsServer(simCardService *service.SimCardService, numeroTelefonicoService *service.NumeroTelefonicoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *SimCardsServer {
	return &SimCardsServer{
		simCardService:          simCardService,
		numeroTelefonicoService: numeroTelefonicoService,
		permissaoMiddleware:     permissaoMiddleware,
	}
}

func (simCardsServer *SimCardsServer) mustEmbedUnimplementedSimCardsServer() {}

// Função para buscar por todos os SimCards existentes
func (simCardsServer *SimCardsServer) FindAllSimCards(context context.Context, req *pb.RequestVazio) (*pb.ListaSimCards, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-simcards")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaSimCards{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCards, erroService := simCardsServer.simCardService.FindAllSimCards(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os SimCards",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaSimCards{SimCards: simCards}, nil
}

// Função para buscar por todos os SimCards disponíveis, ou seja, sem números vinculados
func (simCardsServer *SimCardsServer) FindSimCardsDisponiveis(context context.Context, req *pb.RequestVazio) (*pb.ListaSimCards, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "get-simcards-disponiveis")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaSimCards{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCards, erroService := simCardsServer.simCardService.FindSimCardsDisponiveis(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os SimCards disponíveis",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaSimCards{SimCards: simCards}, nil
}

// Função para buscar por um SimCard pelo ID
func (simCardsServer *SimCardsServer) FindSimCardById(context context.Context, req *pb.RequestId) (*pb.SimCard, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "get-simcard-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCard{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido")
	}

	simCard, erroService := simCardsServer.simCardService.FindSimCardById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um SimCard pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simcard", simCard),
	)

	return simCard, nil
}

// Função para buscar por um SimCard pelo ID
func (simCardsServer *SimCardsServer) FindSimCardByIccid(context context.Context, req *pb.RequestIccid) (*pb.SimCard, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "get-simcard-by-iccid")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCard{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	iccid := req.GetIccid()
	if iccid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ICCID enviado não é válido")
	}

	simCard, erroService := simCardsServer.simCardService.FindSimCardByIccid(context, iccid)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um SimCard pelo ICCID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simcard", simCard),
	)

	return simCard, nil
}

// Função para buscar por um SimCard pelo ID do número telefônico correspondente
func (simCardsServer *SimCardsServer) FindSimCardByTelefoniaNumeroID(context context.Context, req *pb.RequestId) (*pb.SimCard, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "get-simcard-by-telefonianumero-id")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCard{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido")
	}

	simCard, erroService := simCardsServer.simCardService.FindSimCardByTelefoniaNumeroID(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um SimCard pelo id do número telefônico correspondente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simcard", simCard),
	)

	return simCard, nil
}

// Função para buscar por um SimCard pelo ID do número telefônico correspondente
func (simCardsServer *SimCardsServer) FindSimCardByTelefoniaNumero(context context.Context, req *pb.RequestNumero) (*pb.SimCard, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "get-simcard-by-telefonianumero-numero")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCard{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numeroTelefonicoRecebido, erroService := simCardsServer.numeroTelefonicoService.FindNumeroTelefonicoByNumero(context, req.GetNumero())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	simCard, erroService := simCardsServer.simCardService.FindSimCardByTelefoniaNumeroID(context, numeroTelefonicoRecebido.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um SimCard pelo id do número telefônico correspondente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simcard", simCard),
	)

	return simCard, nil
}

// Função para criar um novo SimCard
func (simCardsServer *SimCardsServer) CreateSimCard(context context.Context, req *pb.SimCard) (*pb.SimCard, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-simcard")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCard{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCard, erroService := simCardsServer.simCardService.CreateSimCard(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo SimCard",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simcard", simCard),
	)

	return simCard, nil
}

// Função para criar novos SimCards em lote a partir de um arquivo CSV
func (simCardsServer *SimCardsServer) CreateSimCardCSV(context context.Context, req *pb.ListaSimCards) (*pb.RetornoCSV, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-simcard-csv")
	if retornoMiddleware.Erro != nil {
		return &pb.RetornoCSV{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	var response pb.RetornoCSV
	for i, simCard := range req.GetSimCards() {
		// Converte o índice em string
		iString := strconv.Itoa(i)

		simCardCriado, erroService := simCardsServer.simCardService.CreateSimCard(context, simCard)
		if erroService.Erro != nil {
			logger.Logger.Error("Erro ao criar o SimCard "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
			response.Status = append(response.Status, &pb.CSVStatus{
				Linha:    iString,
				Mensagem: erroService.Erro.Error(),
				Status:   int32(erroService.Status),
			})
			continue
		}

		logger.Logger.Info("Criado um novo SimCard",
			zap.Any("usuario", usuarioSolicitante.Usuario),
			zap.Any("simCard", simCardCriado),
		)

		response.Status = append(response.Status, &pb.CSVStatus{
			Linha:    iString,
			Mensagem: "Sucesso",
			Status:   int32(codes.OK),
		})
	}

	return &response, nil
}

// Função para criar um novo número telefônico
func (simCardsServer *SimCardsServer) VincularSimCard(context context.Context, req *pb.RequestVinculoSimcard) (*pb.SimCard, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "post-vincular-simCard")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCard{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCard, erroService := simCardsServer.simCardService.VincularNumeroTelefonico(context, req.GetSimCardId(), req.GetNumeroTelefonicoId())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Vinculado um número telefônico a um SimCard",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simCard", simCard),
		zap.Any("numero", simCard.GetTelefoniaNumero()),
	)

	return simCard, nil
}

// Função para atualizar um SimCard já existente no banco
func (simCardsServer *SimCardsServer) UpdateSimCard(context context.Context, req *pb.SimCard) (*pb.SimCard, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-simcard")
	if retornoMiddleware.Erro != nil {
		return &pb.SimCard{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	simCardAntigo, erroService := simCardsServer.simCardService.FindSimCardById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	simCard, erroService := simCardsServer.simCardService.UpdateSimCard(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	// TODO Fazer um método separado para mudar o status de um SimCard
	if strings.ToLower(simCard.GetStatus().GetNome()) == "cancelado" {
		mesesQuarentena, err := strconv.Atoi(os.Getenv("PERIODO_QUARENTENA"))
		if err != nil {
			return &pb.SimCard{}, status.Errorf(codes.Internal, err.Error())
		}

		telefoneAntigo := simCard.GetTelefoniaNumero()

		dataCongelado := time.Now().AddDate(0, mesesQuarentena, 0).Format("2006/01/02")
		simCard.GetTelefoniaNumero().CongeladoAte = dataCongelado
		simCard.GetTelefoniaNumero().Utilizavel = false

		simCardsServer.numeroTelefonicoService.UpdateNumeroTelefonico(context, simCard.GetTelefoniaNumero(), telefoneAntigo)
	}

	logger.Logger.Info("Atualizado um SimCard existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("simCardAntigo", simCardAntigo),
		zap.Any("simCardAtualizado", simCard),
	)

	return simCard, nil
}

// Função para deletar um SimCard existente no banco
func (simCardsServer *SimCardsServer) DeleteSimCard(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := simCardsServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-simcard-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido")
	}

	simCard, erroService := simCardsServer.simCardService.FindSimCardById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := simCardsServer.simCardService.DeleteSimCardById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.Aborted, "Não existe SimCard com o ID enviado")
	}

	logger.Logger.Info("Deletado um SimCard",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", simCard),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
