package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
)

type UsuarioJogoRepository interface {
	FindAll(context context.Context) ([]db.TUsuarioJogo, error)
	FindByID(context context.Context, id int32) (db.TUsuarioJogo, error)
	FindByJogo(context context.Context, id int32) ([]db.TUsuarioJogo, error)
	FindByUsuario(context context.Context, id int32) ([]db.TUsuarioJogo, error)
	Create(context context.Context, usuarioJogo db.CreateUsuarioJogoParams) (db.TUsuarioJogo, error)
	Update(context context.Context, usuarioJogo db.UpdateUsuarioJogoParams) (db.TUsuarioJogo, error)
	Delete(context context.Context, id int32) error
}

type usuarioJogoRepository struct {
	*db.Queries
}

func NewUsuarioJogoRepository(queries *db.Queries) UsuarioJogoRepository {
	return &usuarioJogoRepository{Queries: queries}
}

func (usuarioJogoRepository *usuarioJogoRepository) FindAll(context context.Context) ([]db.TUsuarioJogo, error) {
	usuariosJogo, err := usuarioJogoRepository.FindAllUsuarioJogos(context)
	if err != nil {
		return nil, err
	}

	return usuariosJogo, nil
}

func (usuarioJogoRepository *usuarioJogoRepository) FindByID(context context.Context, id int32) (db.TUsuarioJogo, error) {
	usuarioJogo, err := usuarioJogoRepository.FindUsuarioJogoByID(context, id)
	if err != nil {
		return db.TUsuarioJogo{}, err
	}

	return usuarioJogo, nil
}

func (usuarioJogoRepository *usuarioJogoRepository) FindByJogo(context context.Context, perfilId int32) ([]db.TUsuarioJogo, error) {
	usuarioJogo, err := usuarioJogoRepository.FindUsuarioJogoByJogo(context, perfilId)
	if err != nil {
		return []db.TUsuarioJogo{}, err
	}

	return usuarioJogo, nil
}

func (usuarioJogoRepository *usuarioJogoRepository) FindByUsuario(context context.Context, usuarioId int32) ([]db.TUsuarioJogo, error) {
	usosUsuario, err := usuarioJogoRepository.FindUsuarioJogoByUsuario(context, usuarioId)
	if err != nil {
		return []db.TUsuarioJogo{}, err
	}

	return usosUsuario, nil
}

func (usuarioJogoRepository *usuarioJogoRepository) Create(context context.Context, usuarioJogo db.CreateUsuarioJogoParams) (db.TUsuarioJogo, error) {
	usuarioJogoCriado, err := usuarioJogoRepository.CreateUsuarioJogo(context, usuarioJogo)
	if err != nil {
		return db.TUsuarioJogo{}, err
	}

	return usuarioJogoCriado, nil

}

func (usuarioJogoRepository *usuarioJogoRepository) Update(context context.Context, usuarioJogo db.UpdateUsuarioJogoParams) (db.TUsuarioJogo, error) {
	usuarioJogoAtualizado, err := usuarioJogoRepository.UpdateUsuarioJogo(context, usuarioJogo)
	if err != nil {
		return db.TUsuarioJogo{}, err
	}

	return usuarioJogoAtualizado, nil
}

func (usuarioJogoRepository *usuarioJogoRepository) Delete(context context.Context, id int32) error {
	_, err := usuarioJogoRepository.DeleteUsuarioJogoById(context, id)
	return err
}
