package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
)

type FavoritoRepository interface {
	FindByID(context context.Context, id int32) (db.TFavorito, error)
	FindByUsuario(context context.Context, id int32) ([]db.TFavorito, error)
	Create(context context.Context, favorito db.CreateFavoritoParams) (db.TFavorito, error)
	Update(context context.Context, favorito db.UpdateFavoritoParams) (db.TFavorito, error)
	Delete(context context.Context, id int32) (int64, error)
}

type favoritoRepository struct {
	*db.Queries
}

func NewFavoritoRepository(queries *db.Queries) FavoritoRepository {
	return &favoritoRepository{
		Queries: queries,
	}
}

func (favoritoRepository *favoritoRepository) FindByID(context context.Context, id int32) (db.TFavorito, error) {
	favorito, err := favoritoRepository.FindFavoritoByID(context, id)
	if err != nil {
		return db.TFavorito{}, err
	}

	return favorito, nil
}

func (favoritoRepository *favoritoRepository) FindByUsuario(context context.Context, id int32) ([]db.TFavorito, error) {
	favoritos, err := favoritoRepository.FindFavoritoByUsuario(context, id)
	if err != nil {
		return []db.TFavorito{}, err
	}

	return favoritos, nil
}

func (favoritoRepository *favoritoRepository) Create(context context.Context, favorito db.CreateFavoritoParams) (db.TFavorito, error) {
	favoritoCriada, err := favoritoRepository.CreateFavorito(context, favorito)
	if err != nil {
		return db.TFavorito{}, err
	}

	return favoritoCriada, nil
}

func (favoritoRepository *favoritoRepository) Update(context context.Context, favorito db.UpdateFavoritoParams) (db.TFavorito, error) {
	favoritoAtualizada, err := favoritoRepository.UpdateFavorito(context, favorito)
	if err != nil {
		return db.TFavorito{}, err
	}

	return favoritoAtualizada, nil
}

func (favoritoRepository *favoritoRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := favoritoRepository.DeleteFavoritoById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
