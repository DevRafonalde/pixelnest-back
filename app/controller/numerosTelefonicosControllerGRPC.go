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
type NumerosTelefonicosServer struct {
	pb.UnimplementedNumerosTelefonicosServer
	numeroTelefonicoService *service.NumeroTelefonicoService
	permissaoMiddleware     *middlewares.PermissoesMiddleware
}

func NewNumerosTelefonicosServer(numeroTelefonicoService *service.NumeroTelefonicoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *NumerosTelefonicosServer {
	return &NumerosTelefonicosServer{
		numeroTelefonicoService: numeroTelefonicoService,
		permissaoMiddleware:     permissaoMiddleware,
	}
}

func (numeroTelefonicoServer *NumerosTelefonicosServer) mustEmbedUnimplementedNumerosTelefonicosServer() {
}

// Função para buscar por todos os números telefônicos
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindAllNumerosTelefonicos(context context.Context, req *pb.RequestVazio) (*pb.ListaNumerosTelefonicos, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-numerotelefonico")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaNumerosTelefonicos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numerosTelefonicos, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindAllNumeroTelefonicos(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os números telefônicos",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaNumerosTelefonicos{NumerosTelefonicos: numerosTelefonicos}, nil
}

// Função para buscar por um número telefônico pelo ID
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindNumeroTelefonicoById(context context.Context, req *pb.RequestId) (*pb.NumeroTelefonico, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-numerotelefonico-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.NumeroTelefonico{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	numeroTelefonico, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumeroTelefonicoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um número telefônico pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return numeroTelefonico, nil
}

// Função para buscar por um número telefônico pelo número
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindNumeroTelefonicoByNumero(context context.Context, req *pb.RequestNumero) (*pb.NumeroTelefonico, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-numerotelefonico-by-numero")
	if retornoMiddleware.Erro != nil {
		return &pb.NumeroTelefonico{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numeroTelefonico, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumeroTelefonicoByNumero(context, req.GetNumero())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um número telefônico pelo número",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return numeroTelefonico, nil
}

// Função para buscar por números telefônicos disponíveis em certo código de área
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindNumerosTelefonicosDisponiveisByCodArea(context context.Context, req *pb.RequestCodArea) (*pb.ListaNumerosTelefonicos, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-numerotelefonico-disponiveis-cod-area")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaNumerosTelefonicos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numerosTelefonicos, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumerosTelefonicosDisponiveisByCodArea(context, req.GetCodArea())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados números telefônicos disponíveis por código de área",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("codArea", req.GetCodArea()),
	)

	return &pb.ListaNumerosTelefonicos{NumerosTelefonicos: numerosTelefonicos}, nil
}

// Função para buscar números telefônicos disponíveis em certo código de área de certa cidade
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindNumerosTelefonicosDisponiveisByCidade(context context.Context, req *pb.RequestCodIbge) (*pb.ListaNumerosTelefonicos, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-numerotelefonico-disponiveis-cidade")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaNumerosTelefonicos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numerosTelefonicos, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumerosTelefonicosDisponiveisByCidade(context, req.GetCodIbge())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados números telefônicos disponíveis por cidade",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("codIbge", req.GetCodIbge()),
	)

	return &pb.ListaNumerosTelefonicos{NumerosTelefonicos: numerosTelefonicos}, nil
}

// Função para reservar uma quantidade X de números telefônicos disponíveis em certo código de área
func (numeroTelefonicoServer *NumerosTelefonicosServer) ReservarNumerosTelefonicosDisponiveisByCodArea(context context.Context, req *pb.RequestCodAreaReserva) (*pb.ListaNumerosTelefonicos, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-reservar-numerotelefonico-disponiveis-cod-area")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaNumerosTelefonicos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numerosTelefonicos, erroService := numeroTelefonicoServer.numeroTelefonicoService.ReservarNumerosTelefonicosDisponiveisByCodArea(context, req.GetCodArea())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados números telefônicos disponíveis por código de área",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("codArea", req.GetCodArea()),
		zap.Any("quantidade", req.GetQuantidade()),
	)

	return &pb.ListaNumerosTelefonicos{NumerosTelefonicos: numerosTelefonicos}, nil
}

// Função para reservar uma quantidade X de números telefônicos disponíveis em certo código de área de certa cidade
func (numeroTelefonicoServer *NumerosTelefonicosServer) ReservarNumerosTelefonicosDisponiveisByCidade(context context.Context, req *pb.RequestCodIbgeReserva) (*pb.ListaNumerosTelefonicos, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-reservar-numerotelefonico-disponiveis-cidade")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaNumerosTelefonicos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numerosTelefonicos, erroService := numeroTelefonicoServer.numeroTelefonicoService.ReservarNumerosTelefonicosDisponiveisByCidade(context, req.GetCodIbge())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Reservados números telefônicos disponíveis por cidade",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("codIbge", req.GetCodIbge()),
		zap.Any("quantidade", req.GetQuantidade()),
	)

	return &pb.ListaNumerosTelefonicos{NumerosTelefonicos: numerosTelefonicos}, nil
}

// Função para buscar por um número telefônico pelo ID do SimCard correspondente
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindNumeroTelefonicoBySimCardId(context context.Context, req *pb.RequestId) (*pb.NumeroTelefonico, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-numerotelefonico-by-simcardid")
	if retornoMiddleware.Erro != nil {
		return &pb.NumeroTelefonico{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	numeroTelefonico, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumeroTelefonicoBySimCardId(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um número telefônico pelo ID do SimCard correspondente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return numeroTelefonico, nil
}

// Função para buscar por um número telefônico pelo ID do SimCard correspondente
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindNumeroTelefonicoBySimCardIccid(context context.Context, req *pb.RequestIccid) (*pb.NumeroTelefonico, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-numerotelefonico-by-simcardiccid")
	if retornoMiddleware.Erro != nil {
		return &pb.NumeroTelefonico{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numeroTelefonico, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumeroTelefonicoBySimCardIccid(context, req.GetIccid())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um número telefônico pelo ICCID do SimCard correspondente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return numeroTelefonico, nil
}

// Função para buscar por um número telefônico pelo ID do SimCard correspondente
func (numeroTelefonicoServer *NumerosTelefonicosServer) FindNumerosTelefonicosByDocumentoCliente(context context.Context, req *pb.RequestDocumento) (*pb.ListaNumerosTelefonicos, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-numerotelefonico-by-cliente-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaNumerosTelefonicos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numerosTelefonicos, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumerosTelefonicosByDocumentoCliente(context, req.GetDocumento())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um número telefônico pelo ICCID do SimCard correspondente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaNumerosTelefonicos{NumerosTelefonicos: numerosTelefonicos}, nil
}

// Função para criar um novo número telefônico
func (numeroTelefonicoServer *NumerosTelefonicosServer) CreateNumeroTelefonico(context context.Context, req *pb.NumeroTelefonico) (*pb.NumeroTelefonico, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-numerotelefonico")
	if retornoMiddleware.Erro != nil {
		return &pb.NumeroTelefonico{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numeroTelefonico, erroService := numeroTelefonicoServer.numeroTelefonicoService.CreateNumeroTelefonico(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo número telefônico",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", numeroTelefonico),
	)

	return numeroTelefonico, nil
}

// Função para criar números telefônicos em lote a partir de um arquivo csv
func (numeroTelefonicoServer *NumerosTelefonicosServer) CreateNumeroTelefonicoByRange(req *pb.RequestRange, stream pb.NumerosTelefonicos_CreateNumeroTelefonicoByRangeServer) error {
	// Autorização de permissão
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(stream.Context(), "post-create-numerotelefonico-csv")
	if retornoMiddleware.Erro != nil {
		return status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	// Resposta imediata ao cliente antes de iniciar o processo em segundo plano
	err := stream.Send(&pb.RetornoRange{
		Linha:     "inicial",
		Mensagem:  "Processamento iniciado",
		Status:    int32(codes.OK),
		Progresso: 0,
	})
	if err != nil {
		return err
	}

	// Processamento em segundo plano
	go func() {
		// Crie um contexto separado para o processamento em segundo plano
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// total := req.GetRangeFinal() - req.GetRangeInicial()
		for i := req.GetRangeInicial(); i <= req.GetRangeFinal(); i++ {
			// iString := strconv.Itoa(int(i))

			// Cria o objeto Número Telefônico
			numero := pb.NumeroTelefonico{
				CodArea:    req.GetCodArea(),
				Numero:     i,
				Utilizavel: true,
				CodigoCNL:  req.GetCodCNL(),
				ExternalID: req.GetExternalID(),
			}

			// Chama o serviço de criação de número telefônico
			_, erroService := numeroTelefonicoServer.numeroTelefonicoService.CreateNumeroTelefonico(ctx, &numero)
			if erroService.Erro != nil {
				logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
				continue
			}

			logger.Logger.Info("Criado um novo número telefônico",
				zap.Any("usuario", usuarioSolicitante.Usuario),
				zap.Any("numeroTelefonico", &numero),
			)
		}

		// Opcionalmente, você pode registrar quando o processamento em segundo plano termina
		logger.Logger.Info("Processamento de números telefônicos concluído",
			zap.Any("usuario", usuarioSolicitante.Usuario),
			zap.Int32("range_inicial", req.GetRangeInicial()),
			zap.Int32("range_final", req.GetRangeFinal()),
		)
	}()

	// Encerra o stream imediatamente após iniciar o processamento em segundo plano
	return nil
}

// Função para criar números telefônicos em lote a partir de um arquivo csv
func (numeroTelefonicoServer *NumerosTelefonicosServer) CreateNumeroTelefonicoCSV(context context.Context, req *pb.ListaNumerosTelefonicos) (*pb.RetornoCSV, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-numerotelefonico-csv")
	if retornoMiddleware.Erro != nil {
		return &pb.RetornoCSV{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	var response pb.RetornoCSV
	for i, numeroTelefonico := range req.GetNumerosTelefonicos() {
		// Converte o índice em string
		iString := strconv.Itoa(i)

		numeroCriado, erroService := numeroTelefonicoServer.numeroTelefonicoService.CreateNumeroTelefonico(context, numeroTelefonico)
		if erroService.Erro != nil {
			logger.Logger.Error("Erro ao criar o número telefônico "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
			response.Status = append(response.Status, &pb.CSVStatus{
				Linha:    iString,
				Mensagem: erroService.Erro.Error(),
				Status:   int32(erroService.Status),
			})
			continue
		}

		logger.Logger.Info("Criado um novo número telefônico",
			zap.Any("usuario", usuarioSolicitante.Usuario),
			zap.Any("numeroTelefonico", numeroCriado),
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
func (numeroTelefonicoServer *NumerosTelefonicosServer) VincularSimCard(context context.Context, req *pb.RequestVinculoSimcard) (*pb.NumeroTelefonico, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-vincular-simCard")
	if retornoMiddleware.Erro != nil {
		return &pb.NumeroTelefonico{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numeroTelefonico, erroService := numeroTelefonicoServer.numeroTelefonicoService.VincularSimCard(context, req.GetSimCardId(), req.GetNumeroTelefonicoId())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Vinculado um número telefônico a um SimCard",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", numeroTelefonico),
		zap.Any("simCard", numeroTelefonico.GetSimCard()),
	)

	return numeroTelefonico, nil
}

// Função para atualizar um número telefônico existente no banco de dados
func (numeroTelefonicoServer *NumerosTelefonicosServer) UpdateNumeroTelefonico(context context.Context, req *pb.NumeroTelefonico) (*pb.NumeroTelefonico, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-numerotelefonico")
	if retornoMiddleware.Erro != nil {
		return &pb.NumeroTelefonico{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	numeroTelefonicoAntigo, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumeroTelefonicoById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	numeroTelefonicoAtualizado, erroService := numeroTelefonicoServer.numeroTelefonicoService.UpdateNumeroTelefonico(context, req, numeroTelefonicoAntigo)
	if erroService.Erro != nil {
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo número telefônico via CSV",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", numeroTelefonicoAntigo),
		zap.Any("numeroAtualizado", numeroTelefonicoAtualizado),
	)

	return numeroTelefonicoAtualizado, nil
}

// Função para deletar um número telefônico existente no banco de dados
func (numeroTelefonicoServer *NumerosTelefonicosServer) DeleteNumeroTelefonico(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := numeroTelefonicoServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-numerotelefonico-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	numeroTelefonico, erroService := numeroTelefonicoServer.numeroTelefonicoService.FindNumeroTelefonicoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := numeroTelefonicoServer.numeroTelefonicoService.DeleteNumeroTelefonicoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		logger.Logger.Error("Não existe número telefônico com o ID enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe número telefônico com o ID enviado")
	}

	logger.Logger.Info("Deletado um número telefônico",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("numero", numeroTelefonico),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
