package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
)

type UsuarioRepository interface {
	FindAll(context context.Context) ([]db.TUsuario, error)
	FindByID(context context.Context, id int32) (db.TUsuario, error)
	FindByEmail(context context.Context, email string) (db.TUsuario, error)
	Create(context context.Context, usuario db.CreateUsuarioParams) (db.TUsuario, error)
	Update(context context.Context, usuario db.UpdateUsuarioParams) (db.TUsuario, error)
}

type usuarioRepository struct {
	*db.Queries
}

func NewUsuarioRepository(queries *db.Queries) UsuarioRepository {
	return &usuarioRepository{Queries: queries}
}

func (usuarioRepository *usuarioRepository) FindAll(context context.Context) ([]db.TUsuario, error) {
	usuarios, err := usuarioRepository.FindAllUsuarios(context)
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (usuarioRepository *usuarioRepository) FindByID(context context.Context, id int32) (db.TUsuario, error) {
	usuario, err := usuarioRepository.FindUsuarioById(context, id)
	if err != nil {
		return db.TUsuario{}, err
	}

	return usuario, nil
}

func (usuarioRepository *usuarioRepository) FindByEmail(context context.Context, email string) (db.TUsuario, error) {
	usuario, err := usuarioRepository.FindUsuarioByEmail(context, email)
	if err != nil {
		return db.TUsuario{}, err
	}

	return usuario, nil
}

func (usuarioRepository *usuarioRepository) Create(context context.Context, usuario db.CreateUsuarioParams) (db.TUsuario, error) {
	usuarioCriado, err := usuarioRepository.CreateUsuario(context, usuario)
	if err != nil {
		return db.TUsuario{}, err
	}

	return usuarioCriado, nil

}

func (usuarioRepository *usuarioRepository) Update(context context.Context, usuario db.UpdateUsuarioParams) (db.TUsuario, error) {
	usuarioAtualizado, err := usuarioRepository.UpdateUsuario(context, usuario)
	if err != nil {
		return db.TUsuario{}, err
	}

	return usuarioAtualizado, nil
}
