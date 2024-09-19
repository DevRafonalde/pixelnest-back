package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProdutoRepository interface {
	FindAll(context context.Context) ([]db.TProduto, error)
	FindByID(context context.Context, id int32) (db.TProduto, error)
	FindByNome(context context.Context, nome string) ([]db.TProduto, error)
	FindByGenero(context context.Context, genero string) ([]db.TProduto, error)
	Create(context context.Context, produto db.CreateProdutoParams) (db.TProduto, error)
	Update(context context.Context, produto db.UpdateProdutoParams) (db.TProduto, error)
	Delete(context context.Context, id int32) (int64, error)
}

type produtoRepository struct {
	*db.Queries
}

func NewProdutoRepository(queries *db.Queries) ProdutoRepository {
	return &produtoRepository{Queries: queries}
}

func (produtoRepository *produtoRepository) FindAll(context context.Context) ([]db.TProduto, error) {
	produtos, err := produtoRepository.FindAllProdutos(context)
	if err != nil {
		return nil, err
	}

	return produtos, nil
}

func (produtoRepository *produtoRepository) FindByID(context context.Context, id int32) (db.TProduto, error) {
	produto, err := produtoRepository.FindProdutoByID(context, id)
	if err != nil {
		return db.TProduto{}, err
	}

	return produto, nil
}

func (produtoRepository *produtoRepository) FindByNome(context context.Context, nome string) ([]db.TProduto, error) {
	produto, err := produtoRepository.FindProdutoByNome(context, pgtype.Text{String: nome, Valid: true})
	if err != nil {
		return []db.TProduto{}, err
	}

	return produto, nil
}

func (produtoRepository *produtoRepository) FindByGenero(context context.Context, genero string) ([]db.TProduto, error) {
	produto, err := produtoRepository.FindProdutoByGenero(context, pgtype.Text{String: genero, Valid: true})
	if err != nil {
		return []db.TProduto{}, err
	}

	return produto, nil
}

func (produtoRepository *produtoRepository) Create(context context.Context, produto db.CreateProdutoParams) (db.TProduto, error) {
	produtoCriado, err := produtoRepository.CreateProduto(context, produto)
	if err != nil {
		return db.TProduto{}, err
	}

	return produtoCriado, nil
}

func (produtoRepository *produtoRepository) Update(context context.Context, produto db.UpdateProdutoParams) (db.TProduto, error) {
	produtoAtualizado, err := produtoRepository.UpdateProduto(context, produto)
	if err != nil {
		return db.TProduto{}, err
	}

	return produtoAtualizado, nil
}

func (produtoRepository *produtoRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := produtoRepository.DeleteProdutoById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
