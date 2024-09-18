-- name: FindAllUsuarios :many
SELECT * FROM t_usuarios;

-- name: FindUsuarioById :one
SELECT * FROM t_usuarios
WHERE id = $1;

-- name: FindUsuarioByEmail :one
SELECT * FROM t_usuarios
WHERE email = $1;

-- name: CreateUsuario :one
INSERT INTO t_usuarios (nome, email, senha, ativo, token_reset_senha, data_ultima_atualizacao, senha_atualizada)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateUsuario :one
UPDATE t_usuarios
SET nome = $2, email = $3, senha = $4, ativo = $5, token_reset_senha = $6, data_ultima_atualizacao = $7, senha_atualizada = $8
WHERE id = $1
RETURNING *;
