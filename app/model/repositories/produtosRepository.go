package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type SimCardRepository interface {
	FindAll(context context.Context) ([]db.TSimcard, error)
	FindByID(context context.Context, id int32) (db.TSimcard, error)
	FindByIccid(context context.Context, iccid string) (db.TSimcard, error)
	FindByTelefoniaNumeroID(context context.Context, id int32) (db.TSimcard, error)
	Create(context context.Context, simCard db.CreateSimCardParams) (db.TSimcard, error)
	Update(context context.Context, simCard db.UpdateSimCardParams) (db.TSimcard, error)
	Delete(context context.Context, id int32) (int64, error)
}

type simCardRepository struct {
	*db.Queries
}

func NewSimCardRepository(queries *db.Queries) SimCardRepository {
	return &simCardRepository{Queries: queries}
}

func (simCardRepository *simCardRepository) FindAll(context context.Context) ([]db.TSimcard, error) {
	simCards, err := simCardRepository.FindAllSimCards(context)
	if err != nil {
		return nil, err
	}

	return simCards, nil
}

func (simCardRepository *simCardRepository) FindByID(context context.Context, id int32) (db.TSimcard, error) {
	simCard, err := simCardRepository.FindSimCardByID(context, id)
	if err != nil {
		return db.TSimcard{}, err
	}

	return simCard, nil
}

func (simCardRepository *simCardRepository) FindByIccid(context context.Context, iccid string) (db.TSimcard, error) {
	simCard, err := simCardRepository.FindSimCardByIccid(context, iccid)
	if err != nil {
		return db.TSimcard{}, err
	}

	return simCard, nil
}

func (simCardRepository *simCardRepository) FindByTelefoniaNumeroID(context context.Context, id int32) (db.TSimcard, error) {
	simCard, err := simCardRepository.FindSimCardByTelefoniaNumeroID(context, pgtype.Int4{Int32: id, Valid: true})
	if err != nil {
		return db.TSimcard{}, err
	}

	return simCard, nil
}

func (simCardRepository *simCardRepository) Create(context context.Context, simCard db.CreateSimCardParams) (db.TSimcard, error) {
	simCardCriado, err := simCardRepository.CreateSimCard(context, simCard)
	if err != nil {
		return db.TSimcard{}, err
	}

	return simCardCriado, nil
}

func (simCardRepository *simCardRepository) Update(context context.Context, simCard db.UpdateSimCardParams) (db.TSimcard, error) {
	simCardAtualizado, err := simCardRepository.UpdateSimCard(context, simCard)
	if err != nil {
		return db.TSimcard{}, err
	}

	return simCardAtualizado, nil
}

func (simCardRepository *simCardRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := simCardRepository.DeleteSimCardById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
