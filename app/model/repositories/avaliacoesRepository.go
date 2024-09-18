package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
)

type CidadeRepository interface {
	FindAll(context context.Context) ([]db.TCidade, error)
	FindByID(context context.Context, id int32) (db.TCidade, error)
	FindByNome(context context.Context, nome string) (db.TCidade, error)
	FindByUF(context context.Context, uf string) ([]db.TCidade, error)
	FindByCodIbge(context context.Context, codIbge int32) (db.TCidade, error)
	Create(context context.Context, cidade db.CreateCidadeParams) (db.TCidade, error)
	Update(context context.Context, cidade db.UpdateCidadeParams) (db.TCidade, error)
	Delete(context context.Context, id int32) (int64, error)
}

type cidadeRepository struct {
	*db.Queries
}

func NewCidadeRepository(queries *db.Queries) CidadeRepository {
	return &cidadeRepository{
		Queries: queries,
	}
}

func (cidadeRepository *cidadeRepository) FindAll(ctx context.Context) ([]db.TCidade, error) {
	return cidadeRepository.FindAllCidades(ctx)

}

func (cidadeRepository *cidadeRepository) FindByID(context context.Context, id int32) (db.TCidade, error) {
	cidade, err := cidadeRepository.FindCidadeByID(context, id)
	if err != nil {
		return db.TCidade{}, err
	}

	return cidade, nil
}

func (cidadeRepository *cidadeRepository) FindByNome(context context.Context, nome string) (db.TCidade, error) {
	cidade, err := cidadeRepository.FindCidadeByNome(context, nome)
	if err != nil {
		return db.TCidade{}, err
	}

	return cidade, nil
}

func (cidadeRepository *cidadeRepository) FindByUF(context context.Context, uf string) ([]db.TCidade, error) {
	cidades, err := cidadeRepository.FindCidadeByUF(context, uf)
	if err != nil {
		return []db.TCidade{}, err
	}

	return cidades, nil
}

func (cidadeRepository *cidadeRepository) FindByCodIbge(context context.Context, codIbge int32) (db.TCidade, error) {
	cidade, err := cidadeRepository.FindCidadeByCodIbge(context, codIbge)
	if err != nil {
		return db.TCidade{}, err
	}

	return cidade, nil
}

func (cidadeRepository *cidadeRepository) Create(context context.Context, cidade db.CreateCidadeParams) (db.TCidade, error) {
	cidadeCriada, err := cidadeRepository.CreateCidade(context, cidade)
	if err != nil {
		return db.TCidade{}, err
	}

	return cidadeCriada, nil
}

func (cidadeRepository *cidadeRepository) Update(context context.Context, cidade db.UpdateCidadeParams) (db.TCidade, error) {
	cidadeAtualizada, err := cidadeRepository.UpdateCidade(context, cidade)
	if err != nil {
		return db.TCidade{}, err
	}

	return cidadeAtualizada, nil
}

func (cidadeRepository *cidadeRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := cidadeRepository.DeleteCidadeById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
