package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type NumeroTelefonicoRepository interface {
	FindAll(context context.Context) ([]db.TTelefoniaNumero, error)
	FindByID(context context.Context, id int32) (db.TTelefoniaNumero, error)
	FindByNumero(context context.Context, numero int32) ([]db.TTelefoniaNumero, error)
	FindByCodArea(context context.Context, numero int32) ([]db.TTelefoniaNumero, error)
	FindDisponiveis(context context.Context, params db.FindNumerosTelefonicosDisponiveisParams) ([]db.TTelefoniaNumero, error)
	FindBySimCardId(context context.Context, id int32) (db.TTelefoniaNumero, error)
	FindByClienteId(context context.Context, id int32) ([]db.TTelefoniaNumero, error)
	Create(context context.Context, numeroTelefonico db.CreateNumeroTelefonicoParams) (db.TTelefoniaNumero, error)
	Update(context context.Context, numeroTelefonico db.UpdateNumeroTelefonicoParams) (db.TTelefoniaNumero, error)
	Delete(context context.Context, id int32) (int64, error)
}

type numeroTelefonicoRepository struct {
	*db.Queries
}

func NewNumeroTelefonicoRepository(queries *db.Queries) NumeroTelefonicoRepository {
	return &numeroTelefonicoRepository{Queries: queries}
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) FindAll(context context.Context) ([]db.TTelefoniaNumero, error) {
	numerosTelefonicos, err := numeroTelefonicoRepository.FindAllNumerosTelefonicos(context)
	if err != nil {
		return nil, err
	}

	return numerosTelefonicos, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) FindByID(context context.Context, id int32) (db.TTelefoniaNumero, error) {
	numeroTelefonico, err := numeroTelefonicoRepository.FindNumeroTelefonicoByID(context, id)
	if err != nil {
		return db.TTelefoniaNumero{}, err
	}

	return numeroTelefonico, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) FindByNumero(context context.Context, numero int32) ([]db.TTelefoniaNumero, error) {
	numeroTelefonico, err := numeroTelefonicoRepository.FindNumeroTelefonicoByNumero(context, numero)
	if err != nil {
		return []db.TTelefoniaNumero{}, err
	}

	return numeroTelefonico, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) FindByCodArea(context context.Context, codArea int32) ([]db.TTelefoniaNumero, error) {
	numeroTelefonico, err := numeroTelefonicoRepository.FindNumeroTelefonicoByCodArea(context, int16(codArea))
	if err != nil {
		return []db.TTelefoniaNumero{}, err
	}

	return numeroTelefonico, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) FindDisponiveis(context context.Context, params db.FindNumerosTelefonicosDisponiveisParams) ([]db.TTelefoniaNumero, error) {
	numeroTelefonico, err := numeroTelefonicoRepository.FindNumerosTelefonicosDisponiveis(context, params)
	if err != nil {
		return []db.TTelefoniaNumero{}, err
	}

	return numeroTelefonico, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) FindBySimCardId(context context.Context, id int32) (db.TTelefoniaNumero, error) {
	numeroTelefonico, err := numeroTelefonicoRepository.FindNumeroTelefonicoBySimCardId(context, pgtype.Int4{Int32: id, Valid: true})
	if err != nil {
		return db.TTelefoniaNumero{}, err
	}

	return numeroTelefonico, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) FindByClienteId(context context.Context, id int32) ([]db.TTelefoniaNumero, error) {
	numeroTelefonico, err := numeroTelefonicoRepository.FindNumeroTelefonicoByClienteId(context, pgtype.Int4{Int32: id, Valid: true})
	if err != nil {
		return []db.TTelefoniaNumero{}, err
	}

	return numeroTelefonico, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) Create(context context.Context, numeroTelefonico db.CreateNumeroTelefonicoParams) (db.TTelefoniaNumero, error) {
	numeroTelefonicoCriado, err := numeroTelefonicoRepository.CreateNumeroTelefonico(context, numeroTelefonico)
	if err != nil {
		return db.TTelefoniaNumero{}, err
	}

	return numeroTelefonicoCriado, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) Update(context context.Context, numeroTelefonico db.UpdateNumeroTelefonicoParams) (db.TTelefoniaNumero, error) {
	numeroTelefonicoAtualizado, err := numeroTelefonicoRepository.UpdateNumeroTelefonico(context, numeroTelefonico)
	if err != nil {
		return db.TTelefoniaNumero{}, err
	}

	return numeroTelefonicoAtualizado, nil
}

func (numeroTelefonicoRepository *numeroTelefonicoRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := numeroTelefonicoRepository.DeleteNumeroTelefonicoById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
