-- name: FindAllPerfis :many
SELECT * FROM t_perfis;

-- name: FindPerfilByID :one
SELECT * FROM t_perfis WHERE id = $1;

-- name: FindPerfilByNome :one
SELECT * FROM t_perfis WHERE nome = $1;

-- name: FindPerfilByPermissao :many
SELECT p.* 
FROM t_perfis p
JOIN t_perfil_permissao pp ON p.id = pp.perfil_id
WHERE pp.permissao_id = $1;

-- name: CreatePerfil :one
INSERT INTO t_perfis (nome, descricao, ativo, data_ultima_atualizacao)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePerfil :one
UPDATE t_perfis 
SET nome = $1, descricao = $2, ativo = $3, data_ultima_atualizacao = $4
WHERE id = $5
RETURNING *;
