-- name: FindAllPermissoes :many
SELECT * FROM t_permissoes;

-- name: FindPermissaoByID :one
SELECT * FROM t_permissoes WHERE id = $1;

-- name: FindPermissaoByNome :one
SELECT * FROM t_permissoes WHERE nome = $1;

-- name: CreatePermissao :one
INSERT INTO t_permissoes (nome, descricao, ativo, data_ultima_atualizacao)
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: UpdatePermissao :one
UPDATE t_permissoes 
SET nome = $1, descricao = $2, ativo = $3, data_ultima_atualizacao = $4
WHERE id = $5
RETURNING *;
