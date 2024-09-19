package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type AvaliacaoRepository interface {
	FindAll(context context.Context) ([]db.TAvaliaco, error)
	FindByID(context context.Context, id int32) (db.TAvaliaco, error)
	FindByUsuario(context context.Context, id int32) ([]db.TAvaliaco, error)
	FindByProduto(context context.Context, id int32) ([]db.TAvaliaco, error)
	FindByJogo(context context.Context, id int32) ([]db.TAvaliaco, error)
	Create(context context.Context, avaliacao db.CreateAvaliacaoParams) (db.TAvaliaco, error)
	Update(context context.Context, avaliacao db.UpdateAvaliacaoParams) (db.TAvaliaco, error)
	Delete(context context.Context, id int32) (int64, error)
}

type avaliacaoRepository struct {
	*db.Queries
}

func NewAvaliacaoRepository(queries *db.Queries) AvaliacaoRepository {
	return &avaliacaoRepository{
		Queries: queries,
	}
}

func (avaliacaoRepository *avaliacaoRepository) FindAll(ctx context.Context) ([]db.TAvaliaco, error) {
	return avaliacaoRepository.FindAllAvaliacoes(ctx)

}

func (avaliacaoRepository *avaliacaoRepository) FindByID(context context.Context, id int32) (db.TAvaliaco, error) {
	avaliacao, err := avaliacaoRepository.FindAvaliacaoById(context, id)
	if err != nil {
		return db.TAvaliaco{}, err
	}

	return avaliacao, nil
}

func (avaliacaoRepository *avaliacaoRepository) FindByUsuario(context context.Context, id int32) ([]db.TAvaliaco, error) {
	avaliacoes, err := avaliacaoRepository.FindAvaliacaoByUsuario(context, id)
	if err != nil {
		return []db.TAvaliaco{}, err
	}

	return avaliacoes, nil
}

func (avaliacaoRepository *avaliacaoRepository) FindByProduto(context context.Context, id int32) ([]db.TAvaliaco, error) {
	avaliacoes, err := avaliacaoRepository.FindAvaliacaoByProduto(context, pgtype.Int4{Int32: id, Valid: true})
	if err != nil {
		return []db.TAvaliaco{}, err
	}

	return avaliacoes, nil
}

func (avaliacaoRepository *avaliacaoRepository) FindByJogo(context context.Context, id int32) ([]db.TAvaliaco, error) {
	avaliacoes, err := avaliacaoRepository.FindAvaliacaoByJogo(context, pgtype.Int4{Int32: id, Valid: true})
	if err != nil {
		return []db.TAvaliaco{}, err
	}

	return avaliacoes, nil
}

func (avaliacaoRepository *avaliacaoRepository) Create(context context.Context, avaliacao db.CreateAvaliacaoParams) (db.TAvaliaco, error) {
	avaliacaoCriada, err := avaliacaoRepository.CreateAvaliacao(context, avaliacao)
	if err != nil {
		return db.TAvaliaco{}, err
	}

	return avaliacaoCriada, nil
}

func (avaliacaoRepository *avaliacaoRepository) Update(context context.Context, avaliacao db.UpdateAvaliacaoParams) (db.TAvaliaco, error) {
	avaliacaoAtualizada, err := avaliacaoRepository.UpdateAvaliacao(context, avaliacao)
	if err != nil {
		return db.TAvaliaco{}, err
	}

	return avaliacaoAtualizada, nil
}

func (avaliacaoRepository *avaliacaoRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := avaliacaoRepository.DeleteAvaliacaoById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
