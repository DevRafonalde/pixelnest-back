package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
)

type PermissaoRepository interface {
	FindAll(context context.Context) ([]db.TPermisso, error)
	FindByID(context context.Context, id int32) (db.TPermisso, error)
	FindByNome(context context.Context, nome string) (db.TPermisso, error)
	Create(context context.Context, permissao db.CreatePermissaoParams) (db.TPermisso, error)
	Update(context context.Context, permissao db.UpdatePermissaoParams) (db.TPermisso, error)
}

type permissaoRepository struct {
	*db.Queries
}

func NewPermissaoRepository(queries *db.Queries) PermissaoRepository {
	return &permissaoRepository{Queries: queries}
}

func (permissaoRepository *permissaoRepository) FindAll(context context.Context) ([]db.TPermisso, error) {
	permissoes, err := permissaoRepository.FindAllPermissoes(context)
	if err != nil {
		return nil, err
	}

	return permissoes, nil
}

func (permissaoRepository *permissaoRepository) FindByID(context context.Context, id int32) (db.TPermisso, error) {
	permissao, err := permissaoRepository.FindPermissaoByID(context, id)
	if err != nil {
		return db.TPermisso{}, err
	}

	return permissao, nil
}

func (permissaoRepository *permissaoRepository) FindByNome(context context.Context, nome string) (db.TPermisso, error) {
	permissao, err := permissaoRepository.FindPermissaoByNome(context, nome)
	if err != nil {
		return db.TPermisso{}, err
	}

	return permissao, nil
}

func (permissaoRepository *permissaoRepository) Create(context context.Context, permissao db.CreatePermissaoParams) (db.TPermisso, error) {
	permissaoCriada, err := permissaoRepository.CreatePermissao(context, permissao)
	if err != nil {
		return db.TPermisso{}, err
	}

	return permissaoCriada, nil
}

func (permissaoRepository *permissaoRepository) Update(context context.Context, permissao db.UpdatePermissaoParams) (db.TPermisso, error) {
	permissaoAtualizada, err := permissaoRepository.UpdatePermissao(context, permissao)
	if err != nil {
		return db.TPermisso{}, err
	}

	return permissaoAtualizada, nil
}
