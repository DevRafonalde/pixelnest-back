package service

import (
	"context"
	"errors"
	"net/http"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	"pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura do serviço de SimCard
type SimCardService struct {
	simCardRepository          repositories.SimCardRepository
	simCardStatusRepository    repositories.SimCardStatusRepository
	numeroTelefonicoRepository repositories.NumeroTelefonicoRepository
	operadoraRepository        repositories.OperadoraRepository
	simCardTelefoneService     *SimCardTelefoneService
	layout                     string
}

// Função para criar uma nova instância de SimCardService
func NewSimCardService(simCardRepository repositories.SimCardRepository, simCardStatusRepository repositories.SimCardStatusRepository, numeroTelefonicoRepository repositories.NumeroTelefonicoRepository, operadoraRepository repositories.OperadoraRepository, simCardTelefoneService *SimCardTelefoneService) *SimCardService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &SimCardService{
		simCardRepository:          simCardRepository,
		simCardStatusRepository:    simCardStatusRepository,
		numeroTelefonicoRepository: numeroTelefonicoRepository,
		operadoraRepository:        operadoraRepository,
		simCardTelefoneService:     simCardTelefoneService,
		layout:                     "2006-01-02", // Define o formato de data/hora
	}
}

// Função para buscar um SimCard pelo ID
func (simCardService *SimCardService) FindSimCardById(context context.Context, id int32) (*grpc.SimCard, erros.ErroStatus) {
	simCard, err := simCardService.simCardRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum SimCard, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("SimCard não encontrado"),
			}
		}

		return &grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return simCardService.converterParaGRPC(context, simCard)
}

// Função para buscar um SimCard pelo ICCID
func (simCardService *SimCardService) FindSimCardByIccid(context context.Context, iccid string) (*grpc.SimCard, erros.ErroStatus) {
	simCard, err := simCardService.simCardRepository.FindByIccid(context, iccid)
	if err != nil {
		// Caso não seja encontrado nenhum SimCard, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("SimCard não encontrado"),
			}
		}

		return &grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return simCardService.converterParaGRPC(context, simCard)
}

// Função para buscar um SimCard pelo ID do número de telefonia
func (simCardService *SimCardService) FindSimCardByTelefoniaNumeroID(context context.Context, id int32) (*grpc.SimCard, erros.ErroStatus) {
	simCard, err := simCardService.simCardRepository.FindByTelefoniaNumeroID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum SimCard, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("SimCard não encontrado"),
			}
		}

		return &grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return simCardService.converterParaGRPC(context, simCard)
}

// Função para buscar todos os SimCards
func (simCardService *SimCardService) FindAllSimCards(context context.Context) ([]*grpc.SimCard, erros.ErroStatus) {
	simCards, err := simCardService.simCardRepository.FindAll(context)
	if err != nil {
		return []*grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum SimCard, retorna code NotFound
	if len(simCards) == 0 {
		return []*grpc.SimCard{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum SimCard encontrado"),
		}
	}

	pbSimCards := []*grpc.SimCard{}

	for _, simCard := range simCards {
		pbSimCard, err := simCardService.converterParaGRPC(context, simCard)
		if err.Erro != nil {
			return nil, err
		}

		pbSimCards = append(pbSimCards, pbSimCard)
	}

	return pbSimCards, erros.ErroStatus{}
}

// Função para buscar todos os SimCards disponíveis, ou seja, sem números vinculados
func (simCardService *SimCardService) FindSimCardsDisponiveis(context context.Context) ([]*grpc.SimCard, erros.ErroStatus) {
	simCards, err := simCardService.simCardRepository.FindAll(context)
	if err != nil {
		return []*grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum SimCard, retorna code NotFound
	if len(simCards) == 0 {
		return []*grpc.SimCard{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum SimCard encontrado"),
		}
	}

	pbSimCards := []*grpc.SimCard{}

	for _, simCard := range simCards {
		pbSimCard, err := simCardService.converterParaGRPC(context, simCard)
		if err.Erro != nil {
			return nil, err
		}

		if pbSimCard.GetStatus().GetNome() == "Disponível" {
			pbSimCards = append(pbSimCards, pbSimCard)
		}
	}

	// Caso nenhum SimCard esteja disponível, retorna code NotFound
	if len(simCards) == 0 {
		return []*grpc.SimCard{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum SimCard disponível encontrado"),
		}
	}

	return pbSimCards, erros.ErroStatus{}
}

// Função para criar um novo SimCard
func (simCardService *SimCardService) CreateSimCard(context context.Context, simCardRecebido *grpc.SimCard) (*grpc.SimCard, erros.ErroStatus) {
	// Faz a conversão do ID int32 que vem do modelo de comunicação (grpc) para int4 do postgres utilizado no modelo de gravação no banco (sqlc)
	var telefoniaNumero pgtype.Int4
	if simCardRecebido.GetTelefoniaNumero().GetID() != 0 {
		telefoniaNumero = pgtype.Int4{Int32: simCardRecebido.GetTelefoniaNumero().GetID(), Valid: true}
	} else {
		telefoniaNumero = pgtype.Int4{Int32: simCardRecebido.GetTelefoniaNumero().GetID(), Valid: false}
	}

	simCardExistenteUsandoNumero, erroBuscaConflito := simCardService.FindSimCardByTelefoniaNumeroID(context, simCardRecebido.GetTelefoniaNumero().GetID())
	if erroBuscaConflito.Erro == nil {
		if simCardExistenteUsandoNumero.GetStatus().GetNome() == "Ativo" {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe um SimCard ativo vinculado a esse número telefônico"),
			}
		}
	}

	// Cria o objeto CreateSimCardParams gerado pelo sqlc para gravação no banco de dados
	simCardCreate := db.CreateSimCardParams{
		Iccid:             simCardRecebido.GetICCID(),
		Imsi:              simCardRecebido.GetIMSI(),
		Pin:               simCardRecebido.GetPIN(),
		Puk:               simCardRecebido.GetPUK(),
		Ki:                simCardRecebido.GetKI(),
		Opc:               simCardRecebido.GetOPC(),
		StatusID:          int16(simCardRecebido.GetStatus().GetID()),
		TelefoniaNumeroID: telefoniaNumero,
		DataCriacao:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
		DataStatus:        pgtype.Timestamptz{Time: time.Now(), Valid: true},
		AtualizadoEm:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Puk2:              pgtype.Text{String: simCardRecebido.GetPUK2(), Valid: true},
		Pin2:              pgtype.Text{String: simCardRecebido.GetPIN2(), Valid: true},
	}

	simCardCriado, err := simCardService.simCardRepository.Create(context, simCardCreate)
	if err != nil {
		// Caso o erro retornado tenha a chave explicitada abaixo, retorna um erro informando que o estado de SimCard não existe
		// É correto considerar apenas essa possibilidade visto que a coluna de estado_id não é única
		// Permitindo que o mesmo estado possa ser usado várias vezes
		// Ou seja, qualquer erro se referenciando à chave estrangeira, significa que aconteceu pois ela não foi encontrada
		if strings.Contains(err.Error(), "fk_t_simcard_estado_sim_cards") {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.InvalidArgument,
				Erro:   errors.New("Não existe estado de SimCard com o id enviado"),
			}
		}

		return &grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return simCardService.converterParaGRPC(context, simCardCriado)
}

// Função para vincular um SimCard a um número telefônico
func (simCardService *SimCardService) VincularNumeroTelefonico(context context.Context, idSimCard int32, idNumeroTelefonico int32) (*grpc.SimCard, erros.ErroStatus) {
	simCard, _, err := simCardService.simCardTelefoneService.Vincular(context, idSimCard, idNumeroTelefonico)
	if err.Erro != nil {
		return &grpc.SimCard{}, err
	}

	return simCardService.converterParaGRPC(context, simCard)
}

// Função para atualizar um SimCard existente
func (simCardService *SimCardService) UpdateSimCard(context context.Context, simCardRecebido *grpc.SimCard) (*grpc.SimCard, erros.ErroStatus) {
	simCardBanco, err := simCardService.simCardRepository.FindByID(context, simCardRecebido.ID)
	if err != nil {
		return &grpc.SimCard{}, erros.ErroStatus{
			Status: http.StatusBadRequest,
			Erro:   err,
		}
	}

	// Faz a conversão das datas string que vêm do modelo de comunicação (grpc) para datas Time utilizadas no modelo de gravação no banco (sqlc)
	var dataStatus time.Time
	if simCardBanco.StatusID != int16(simCardRecebido.GetStatus().GetID()) {
		dataStatus, err = time.Parse("2006/01/02", time.Now().Format("2006/01/02"))
		if err != nil {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	} else {
		dataStatus = simCardBanco.DataStatus.Time
	}

	atualizadoEm, err := time.Parse("2006/01/02", time.Now().Format("2006/01/02"))
	if err != nil {
		return &grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Faz a conversão do ID int32 que vem do modelo de comunicação (grpc) para int4 do postgres utilizado no modelo de gravação no banco (sqlc)
	var telefoniaNumero pgtype.Int4
	if simCardRecebido.GetTelefoniaNumero().GetID() != 0 {
		telefoniaNumero = pgtype.Int4{Int32: simCardRecebido.GetTelefoniaNumero().GetID(), Valid: true}
	} else {
		telefoniaNumero = pgtype.Int4{Int32: simCardRecebido.GetTelefoniaNumero().GetID(), Valid: false}
	}

	// Cria o objeto UpdateSimCardParams gerado pelo sqlc para gravação no banco de dados
	simCardUpdate := db.UpdateSimCardParams{
		Iccid:             simCardRecebido.ICCID,
		Imsi:              simCardRecebido.IMSI,
		Pin:               simCardRecebido.PIN,
		Puk:               simCardRecebido.PUK,
		Ki:                simCardRecebido.KI,
		Opc:               simCardRecebido.OPC,
		StatusID:          int16(simCardRecebido.GetStatus().GetID()),
		TelefoniaNumeroID: telefoniaNumero,
		DataCriacao:       simCardBanco.DataCriacao,
		DataStatus:        pgtype.Timestamptz{Time: dataStatus, Valid: true},
		AtualizadoEm:      pgtype.Timestamptz{Time: atualizadoEm, Valid: true},
		Puk2:              pgtype.Text{String: simCardRecebido.PUK2, Valid: true},
		Pin2:              pgtype.Text{String: simCardRecebido.PIN2, Valid: true},
		ID:                simCardRecebido.GetID(),
	}

	// Salva o SimCard atualizado no repositório
	simCardAtualizado, erroSalvamento := simCardService.simCardRepository.Update(context, simCardUpdate)
	if erroSalvamento != nil {
		// Mesmo caso do create
		if strings.Contains(erroSalvamento.Error(), "fk_t_simcard_estado_sim_cards") {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.InvalidArgument,
				Erro:   errors.New("Não existe estado de SimCard com o id enviado"),
			}
		}

		return &grpc.SimCard{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroSalvamento,
		}
	}

	return simCardService.converterParaGRPC(context, simCardAtualizado)
}

// Função para deletar um SimCard pelo ID
func (simCardService *SimCardService) DeleteSimCardById(context context.Context, id int32) (bool, erros.ErroStatus) {
	deletados, err := simCardService.simCardRepository.Delete(context, id)
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhuma cidade foi deletada, pois não existia
	if deletados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Status de SimCard não encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}

// Essa é uma função privada do serviço para encapsular a montagem do objeto de retorno para o server/controller
func (simCardService *SimCardService) converterParaGRPC(context context.Context, simCard db.TSimcard) (*grpc.SimCard, erros.ErroStatus) {
	// Faz a montagem do objeto estado de SimCard
	estado, err := simCardService.simCardStatusRepository.FindByID(context, int32(simCard.StatusID))
	if err != nil {
		if err.Error() == "no rows in result set" {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: http.StatusBadRequest,
				Erro:   errors.New("Status de SimCard correspondente não encontrado"),
			}
		} else {
			return &grpc.SimCard{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	}

	// Aqui dentro desse grande if, que apenas é executado caso o SimCard esteja vinculado à um número telefônico,
	// É feita basicamente a mesma operação da função `montarObjNumeroCompleto` do serviço de números telefônicos,
	// Com exceção do comentário explicitado mais abaixo
	numeroTelefonico := new(db.TTelefoniaNumero)
	operadoraIn := new(grpc.Operadora)
	operadoraOut := new(grpc.Operadora)
	if simCard.TelefoniaNumeroID.Int32 != 0 {
		tNumeroTelefonico, err := simCardService.numeroTelefonicoRepository.FindByID(context, simCard.TelefoniaNumeroID.Int32)
		if err != nil {
			if err.Error() == "no rows in result set" {
				return &grpc.SimCard{}, erros.ErroStatus{
					Status: http.StatusBadRequest,
					Erro:   errors.New("Número telefônico correspondente não encontrado"),
				}
			} else {
				return &grpc.SimCard{}, erros.ErroStatus{
					Status: codes.Internal,
					Erro:   err,
				}
			}
		}
		numeroTelefonico = &tNumeroTelefonico

		if numeroTelefonico.PortadoInOperadoraID.Int32 != 0 {
			tOperadoraIn, err := simCardService.operadoraRepository.FindByID(context, numeroTelefonico.PortadoInOperadoraID.Int32)
			if err != nil {
				return &grpc.SimCard{}, erros.ErroStatus{
					Status: codes.Internal,
					Erro:   err,
				}
			}
			operadoraIn = helpers.TOperadoraToPb(tOperadoraIn)
		}

		if numeroTelefonico.PortadoOutOperadoraID.Int32 != 0 {
			tOperadoraOut, err := simCardService.operadoraRepository.FindByID(context, numeroTelefonico.PortadoInOperadoraID.Int32)
			if err != nil {
				return &grpc.SimCard{}, erros.ErroStatus{
					Status: codes.Internal,
					Erro:   err,
				}
			}
			operadoraOut = helpers.TOperadoraToPb(tOperadoraOut)
		}
	}

	// Aqui, ao transformar o número telefônico para o modelo de retorno, é passado um valor nulo no último parâmetro
	// Esse valor seria o SimCard correspondente, porém é justamente isso que estou montando aqui
	// Passar esse valor para a conversão faria com que isso se tornasse uma função recursiva e entraria em um loop infitino
	return helpers.TSimCardToPb(simCard, helpers.TStatusSimCardToPb(estado), helpers.TTelefoniaNumeroToPb(*numeroTelefonico, operadoraIn, operadoraOut, nil)), erros.ErroStatus{}
}
