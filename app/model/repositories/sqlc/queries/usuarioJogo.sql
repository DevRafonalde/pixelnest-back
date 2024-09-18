-- name: FindAllUsuarioJogos :many
SELECT * FROM t_usuario_jogo;

-- name: FindUsuarioJogoByID :one
SELECT * FROM t_usuario_jogo WHERE id = $1;

-- name: FindUsuarioJogoByJogo :many
SELECT * FROM t_usuario_jogo WHERE jogo_id = $1;

-- name: FindUsuarioJogoByUsuario :many
SELECT * FROM t_usuario_jogo WHERE usuario_id = $1;

-- name: CreateUsuarioJogo :one
INSERT INTO t_usuario_jogo (usuario_id, jogo_id, data_hora)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUsuarioJogo :one
UPDATE t_usuario_jogo 
SET usuario_id = $1, jogo_id = $2, data_hora = $3
WHERE id = $4
RETURNING *;

-- name: DeleteUsuarioJogoById :execrows
DELETE FROM t_usuario_jogo 
WHERE id = $1 
RETURNING *;
