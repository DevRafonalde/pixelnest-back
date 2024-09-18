package service

import (
	"context"
	"errors"
	"pixelnest/app/model/erros"
	"pixelnest/app/model/repositories"
	"pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura de serviço para gerenciar operações relacionadas às cidades
type SimCardTelefoneService struct {
	simCardRepository          repositories.SimCardRepository
	numeroTelefonicoRepository repositories.NumeroTelefonicoRepository
}

// Função para criar uma nova instância de SimCardTelefoneService com o repositório necessário
func NewSimCardTelefoneService(simCardRepository repositories.SimCardRepository, numeroTelefonicoRepository repositories.NumeroTelefonicoRepository) *SimCardTelefoneService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &SimCardTelefoneService{
		simCardRepository:          simCardRepository,
		numeroTelefonicoRepository: numeroTelefonicoRepository,
	}
}

func (simCardTelefoneService *SimCardTelefoneService) Vincular(context context.Context, idSimCard int32, idNumeroTelefonico int32) (repositoryIMPL.TSimcard, repositoryIMPL.TTelefoniaNumero, erros.ErroStatus) {
	// Busca o número telefônico no repositório pelo ID
	numeroTelefonico, err := simCardTelefoneService.numeroTelefonicoRepository.FindByID(context, idNumeroTelefonico)
	if err != nil {
		// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return db.TSimcard{}, db.TTelefoniaNumero{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Número telefônico não encontrado"),
			}
		}

		return db.TSimcard{}, db.TTelefoniaNumero{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	numeroTelefonicoUpdate := db.UpdateNumeroTelefonicoParams{
		CodArea:               numeroTelefonico.CodArea,
		Numero:                numeroTelefonico.Numero,
		Utilizavel:            numeroTelefonico.Utilizavel,
		PortadoIn:             numeroTelefonico.PortadoIn,
		PortadoInOperadoraID:  numeroTelefonico.PortadoInOperadoraID,
		PortadoInDate:         numeroTelefonico.PortadoInDate,
		CodigoCnl:             numeroTelefonico.CodigoCnl,
		CongeladoAte:          numeroTelefonico.CongeladoAte,
		ExternalID:            numeroTelefonico.ExternalID,
		PortadoOut:            numeroTelefonico.PortadoOut,
		PortadoOutOperadoraID: numeroTelefonico.PortadoOutOperadoraID,
		PortadoOutDate:        numeroTelefonico.PortadoOutDate,
		DataCriacao:           numeroTelefonico.DataCriacao,
		SimCardID:             pgtype.Int4{Int32: idSimCard, Valid: true},
		ID:                    numeroTelefonico.ID,
	}

	// Salva o número telefônico atualizado no repositório
	numeroTelefonicoAtualizado, err := simCardTelefoneService.numeroTelefonicoRepository.Update(context, numeroTelefonicoUpdate)
	if err != nil {
		return db.TSimcard{}, db.TTelefoniaNumero{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Busca o SimCard no repositório pelo ID
	simCard, err := simCardTelefoneService.simCardRepository.FindByID(context, idSimCard)
	if err != nil {
		// Caso não seja encontrado nenhum SimCard, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return db.TSimcard{}, db.TTelefoniaNumero{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("SimCard não encontrado"),
			}
		}

		return db.TSimcard{}, db.TTelefoniaNumero{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	atualizadoEm, err := time.Parse("2006/01/02", time.Now().Format("2006/01/02"))
	if err != nil {
		return db.TSimcard{}, db.TTelefoniaNumero{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	simCardUpdate := db.UpdateSimCardParams{
		Iccid:             simCard.Iccid,
		Imsi:              simCard.Imsi,
		Pin:               simCard.Pin,
		Puk:               simCard.Puk,
		Ki:                simCard.Ki,
		Opc:               simCard.Opc,
		StatusID:          simCard.StatusID,
		TelefoniaNumeroID: pgtype.Int4{Int32: idNumeroTelefonico, Valid: true},
		DataCriacao:       simCard.DataCriacao,
		DataStatus:        simCard.DataStatus,
		AtualizadoEm:      pgtype.Timestamptz{Time: atualizadoEm, Valid: true},
		Puk2:              simCard.Puk2,
		Pin2:              simCard.Pin2,
		ID:                simCard.ID,
	}

	// Salva o SimCard atualizado no repositório
	simCardAtualizado, erroSalvamento := simCardTelefoneService.simCardRepository.Update(context, simCardUpdate)
	if erroSalvamento != nil {
		return db.TSimcard{}, db.TTelefoniaNumero{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroSalvamento,
		}
	}

	return simCardAtualizado, numeroTelefonicoAtualizado, erros.ErroStatus{}
}
