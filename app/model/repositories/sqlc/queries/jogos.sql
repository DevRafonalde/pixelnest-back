-- name: FindAllJogos :many
SELECT * FROM t_jogos;

-- name: FindJogoByID :one
SELECT * FROM t_jogos WHERE id = $1;

-- name: FindJogoByNome :many
SELECT * FROM t_jogos WHERE nome ILIKE $1;

-- name: FindJogoByGenero :many
SELECT * FROM t_jogos WHERE genero ILIKE $1;

-- name: CreateJogo :one
INSERT INTO t_jogos (nome, sinopse, genero)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateJogo :one
UPDATE t_jogos 
SET nome = $1, sinopse = $2, genero = $3
WHERE id = $4
RETURNING *;

-- name: DeleteJogoById :execrows
DELETE FROM t_jogos WHERE id = $1 
RETURNING id;
