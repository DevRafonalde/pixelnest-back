package service

import (
	"context"
	"errors"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	"pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura do serviço de SimCardStatus
type SimCardStatusService struct {
	simCardStatusRepository repositories.SimCardStatusRepository
}

// Função para criar uma nova instância de SimCardStatusService
func NewSimCardStatusService(simCardStatusRepository repositories.SimCardStatusRepository) *SimCardStatusService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &SimCardStatusService{
		simCardStatusRepository: simCardStatusRepository,
	}
}

// Função para buscar um SimCardStatus pelo ID
func (simCardStatusService SimCardStatusService) FindSimCardStatusById(context context.Context, id int32) (*grpc.SimCardStatus, erros.ErroStatus) {
	simCardStatus, err := simCardStatusService.simCardStatusRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum estado de SimCard, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.SimCardStatus{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Status de SimCard não encontrado"),
			}
		}

		return &grpc.SimCardStatus{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TStatusSimCardToPb(simCardStatus), erros.ErroStatus{}
}

// Função para buscar um SimCardStatus pelo nome do estado
func (simCardStatusService SimCardStatusService) FindSimCardStatusByStatus(context context.Context, estado string) (*grpc.SimCardStatus, erros.ErroStatus) {
	simCardStatus, err := simCardStatusService.simCardStatusRepository.FindByNome(context, estado)
	if err != nil {
		// Caso não seja encontrado nenhum estado de SimCard, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.SimCardStatus{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Status de SimCard não encontrado"),
			}
		}

		return &grpc.SimCardStatus{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TStatusSimCardToPb(simCardStatus), erros.ErroStatus{}
}

// Função para buscar todos os SimCardStatus
func (simCardStatusService SimCardStatusService) FindAllSimCardStatus(context context.Context) ([]*grpc.SimCardStatus, erros.ErroStatus) {
	simCardStatus, err := simCardStatusService.simCardStatusRepository.FindAll(context)
	if err != nil {
		return []*grpc.SimCardStatus{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum estado de SimCard, retorna code NotFound
	if len(simCardStatus) == 0 {
		return []*grpc.SimCardStatus{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum estado de SimCard encontrado"),
		}
	}

	var pbStatus []*grpc.SimCardStatus
	for _, estado := range simCardStatus {
		pbStatus = append(pbStatus, helpers.TStatusSimCardToPb(estado))
	}

	return pbStatus, erros.ErroStatus{}
}

// Função para criar um novo SimCardStatus
func (simCardStatusService SimCardStatusService) CreateSimCardStatus(context context.Context, simCardStatus *grpc.SimCardStatus) (*grpc.SimCardStatus, erros.ErroStatus) {
	// Busca um estado de SimCard pelo estado(nome) enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	_, erroBuscaPreExistente := simCardStatusService.FindSimCardStatusByStatus(context, simCardStatus.GetNome())
	if erroBuscaPreExistente.Erro == nil {
		return &grpc.SimCardStatus{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Já existe estado de SimCard com esse nome"),
		}
	}

	// Cria o objeto CreateSimCardStatusParams gerado pelo sqlc para gravação no banco de dados
	simCardStatusCreate := db.CreateSimCardStatusParams{
		Nome:      simCardStatus.GetNome(),
		Descricao: pgtype.Text{String: simCardStatus.GetDescricao(), Valid: true},
	}

	// Cria o estado de SimCard no repositório
	simCardStatusCriado, err := simCardStatusService.simCardStatusRepository.Create(context, simCardStatusCreate)
	if err != nil {
		return &grpc.SimCardStatus{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TStatusSimCardToPb(simCardStatusCriado), erros.ErroStatus{}
}

// Função para atualizar um SimCardStatus existente
func (simCardStatusService SimCardStatusService) UpdateSimCardStatus(context context.Context, simCardStatusRecebido *grpc.SimCardStatus) (*grpc.SimCardStatus, erros.ErroStatus) {
	simCardStatusBanco, erroBuscaPreExistente := simCardStatusService.FindSimCardStatusById(context, simCardStatusRecebido.GetID())
	if erroBuscaPreExistente.Erro != nil {
		return &grpc.SimCardStatus{}, erroBuscaPreExistente
	}

	// Verifica se o estado(nome) foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if simCardStatusBanco.GetNome() != simCardStatusRecebido.GetNome() {
		_, erroBuscaPreExistente := simCardStatusService.FindSimCardStatusByStatus(context, simCardStatusRecebido.GetNome())
		if erroBuscaPreExistente.Erro == nil {
			return &grpc.SimCardStatus{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe estado de SimCard com esse nome"),
			}
		}
	}

	// Cria o objeto UpdateSimCardStatusParams gerado pelo sqlc para gravação no banco de dados
	simCardStatusUpdate := db.UpdateSimCardStatusParams{
		Nome:      simCardStatusRecebido.GetNome(),
		Descricao: pgtype.Text{String: simCardStatusRecebido.GetDescricao(), Valid: true},
		ID:        simCardStatusRecebido.GetID(),
	}

	// Salva o estado de SimCard atualizada no repositório
	simCardStatusAtualizado, erroSalvamento := simCardStatusService.simCardStatusRepository.Update(context, simCardStatusUpdate)
	if erroSalvamento != nil {
		return &grpc.SimCardStatus{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroSalvamento,
		}
	}
	return helpers.TStatusSimCardToPb(simCardStatusAtualizado), erros.ErroStatus{}
}

// Função para deletar um SimCardStatus pelo ID
func (simCardStatusService SimCardStatusService) DeleteSimCardStatusById(context context.Context, id int32) (bool, erros.ErroStatus) {
	deletados, err := simCardStatusService.simCardStatusRepository.Delete(context, id)
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhum estado de SimCard foi deletado, pois não existia
	if deletados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Status de SimCard não encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}
