-- name: FindAllParametros :many
SELECT * FROM t_parametros;

-- name: FindParametroByNome :one
SELECT * FROM t_parametros WHERE nome = $1;

-- name: FindParametroById :one
SELECT * FROM t_parametros WHERE id = $1;

-- name: CreateParametro :one
INSERT INTO t_parametros (nome, valor, descricao) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateParametro :one
UPDATE t_parametros 
SET nome = $1, valor = $2, descricao = $3
WHERE id = $4
RETURNING *;

-- name: DeleteParametroById :execrows
DELETE FROM t_parametros WHERE id = $1 
RETURNING id;
