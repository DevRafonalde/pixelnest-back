package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type JogoRepository interface {
	FindAll(context context.Context) ([]db.TJogo, error)
	FindByID(context context.Context, id int32) (db.TJogo, error)
	FindByNome(context context.Context, nome string) ([]db.TJogo, error)
	FindByGenero(context context.Context, genero string) ([]db.TJogo, error)
	Create(context context.Context, jogo db.CreateJogoParams) (db.TJogo, error)
	Update(context context.Context, jogo db.UpdateJogoParams) (db.TJogo, error)
	Delete(context context.Context, id int32) (int64, error)
}

type jogoRepository struct {
	*db.Queries
}

func NewJogoRepository(queries *db.Queries) JogoRepository {
	return &jogoRepository{Queries: queries}
}

func (jogoRepository *jogoRepository) FindAll(context context.Context) ([]db.TJogo, error) {
	numerosTelefonicos, err := jogoRepository.FindAllJogos(context)
	if err != nil {
		return nil, err
	}

	return numerosTelefonicos, nil
}

func (jogoRepository *jogoRepository) FindByID(context context.Context, id int32) (db.TJogo, error) {
	jogo, err := jogoRepository.FindJogoByID(context, id)
	if err != nil {
		return db.TJogo{}, err
	}

	return jogo, nil
}

func (jogoRepository *jogoRepository) FindByNome(context context.Context, nome string) ([]db.TJogo, error) {
	jogo, err := jogoRepository.FindJogoByNome(context, pgtype.Text{String: nome, Valid: true})
	if err != nil {
		return []db.TJogo{}, err
	}

	return jogo, nil
}

func (jogoRepository *jogoRepository) FindByGenero(context context.Context, genero string) ([]db.TJogo, error) {
	jogo, err := jogoRepository.FindJogoByGenero(context, pgtype.Text{String: genero, Valid: true})
	if err != nil {
		return []db.TJogo{}, err
	}

	return jogo, nil
}

func (jogoRepository *jogoRepository) Create(context context.Context, jogo db.CreateJogoParams) (db.TJogo, error) {
	jogoCriado, err := jogoRepository.CreateJogo(context, jogo)
	if err != nil {
		return db.TJogo{}, err
	}

	return jogoCriado, nil
}

func (jogoRepository *jogoRepository) Update(context context.Context, jogo db.UpdateJogoParams) (db.TJogo, error) {
	jogoAtualizado, err := jogoRepository.UpdateJogo(context, jogo)
	if err != nil {
		return db.TJogo{}, err
	}

	return jogoAtualizado, nil
}

func (jogoRepository *jogoRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := jogoRepository.DeleteJogoById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
