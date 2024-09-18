package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
)

type PerfilRepository interface {
	FindAll(context context.Context) ([]db.TPerfi, error)
	FindByID(context context.Context, id int32) (db.TPerfi, error)
	FindByNome(context context.Context, nome string) (db.TPerfi, error)
	Create(context context.Context, perfil db.CreatePerfilParams) (db.TPerfi, error)
	Update(context context.Context, perfil db.UpdatePerfilParams) (db.TPerfi, error)
}

type perfilRepository struct {
	*db.Queries
}

func NewPerfilRepository(queries *db.Queries) PerfilRepository {
	return &perfilRepository{Queries: queries}
}

func (perfilRepository *perfilRepository) FindAll(context context.Context) ([]db.TPerfi, error) {
	perfis, err := perfilRepository.FindAllPerfis(context)
	if err != nil {
		return nil, err
	}

	return perfis, nil
}

func (perfilRepository *perfilRepository) FindByID(context context.Context, id int32) (db.TPerfi, error) {
	perfil, err := perfilRepository.FindPerfilByID(context, id)
	if err != nil {
		return db.TPerfi{}, err
	}

	return perfil, nil
}

func (perfilRepository *perfilRepository) FindByNome(context context.Context, nome string) (db.TPerfi, error) {
	perfil, err := perfilRepository.FindPerfilByNome(context, nome)
	if err != nil {
		return db.TPerfi{}, err
	}

	return perfil, nil
}

func (perfilRepository *perfilRepository) Create(context context.Context, perfil db.CreatePerfilParams) (db.TPerfi, error) {
	perfilCriado, err := perfilRepository.CreatePerfil(context, perfil)
	if err != nil {
		return db.TPerfi{}, err
	}

	return perfilCriado, nil

}

func (perfilRepository *perfilRepository) Update(context context.Context, perfil db.UpdatePerfilParams) (db.TPerfi, error) {
	perfilAtualizado, err := perfilRepository.UpdatePerfil(context, perfil)
	if err != nil {
		return db.TPerfi{}, err
	}

	return perfilAtualizado, nil
}
