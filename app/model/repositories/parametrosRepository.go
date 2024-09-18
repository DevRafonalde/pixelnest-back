package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type ParametroRepository interface {
	FindAll(context context.Context) ([]db.TParametro, error)
	FindByNome(context context.Context, nome string) (db.TParametro, error)
	FindById(context context.Context, id int32) (db.TParametro, error)
	Create(context context.Context, parametro db.CreateParametroParams) (db.TParametro, error)
	Update(context context.Context, parametro db.UpdateParametroParams) (db.TParametro, error)
	Delete(context context.Context, id int32) (int64, error)
}

type parametroRepository struct {
	*db.Queries
}

func NewParametroRepository(queries *db.Queries) ParametroRepository {
	return &parametroRepository{Queries: queries}
}

func (parametroRepository *parametroRepository) FindAll(context context.Context) ([]db.TParametro, error) {
	parametros, err := parametroRepository.FindAllParametros(context)
	if err != nil {
		return nil, err
	}

	return parametros, nil
}

func (parametroRepository *parametroRepository) FindByNome(context context.Context, nome string) (db.TParametro, error) {
	parametro, err := parametroRepository.FindParametroByNome(context, pgtype.Text{String: nome, Valid: true})
	if err != nil {
		return db.TParametro{}, err
	}

	return parametro, nil
}

func (parametroRepository *parametroRepository) FindById(context context.Context, id int32) (db.TParametro, error) {
	parametro, err := parametroRepository.FindParametroById(context, id)
	if err != nil {
		return db.TParametro{}, err
	}

	return parametro, nil
}

func (parametroRepository *parametroRepository) Create(context context.Context, parametro db.CreateParametroParams) (db.TParametro, error) {
	parametroCriada, err := parametroRepository.CreateParametro(context, parametro)
	if err != nil {
		return db.TParametro{}, err
	}

	return parametroCriada, nil
}

func (parametroRepository *parametroRepository) Update(context context.Context, parametro db.UpdateParametroParams) (db.TParametro, error) {
	parametroAtualizada, err := parametroRepository.UpdateParametro(context, parametro)
	if err != nil {
		return db.TParametro{}, err
	}

	return parametroAtualizada, nil
}

func (parametroRepository *parametroRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfedatas, err := parametroRepository.DeleteParametroById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfedatas, nil
}
