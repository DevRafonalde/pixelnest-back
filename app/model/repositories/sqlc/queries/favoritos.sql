-- name: FindFavoritoByID :one
SELECT * FROM t_favoritos WHERE id = $1;

-- name: FindFavoritoByUsuario :one
SELECT * FROM t_favoritos WHERE usuario_id = $1;

-- name: CreateFavorito :one
INSERT INTO t_favoritos (usuario_id, produto_id, jogo_id) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateFavorito :one
UPDATE t_favoritos 
SET usuario_id = $1, produto_id = $2, jogo_id = $3
WHERE id = $4
RETURNING *;

-- name: DeleteFavoritoById :execrows
DELETE FROM t_favoritos WHERE id = $1 
RETURNING id;
