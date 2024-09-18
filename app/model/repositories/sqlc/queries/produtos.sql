-- name: FindAllProdutos :many
SELECT * FROM t_produtos;

-- name: FindProdutoByID :one
SELECT * FROM t_produtos WHERE id = $1;

-- name: FindProdutoByNome :many
SELECT * FROM t_produtos WHERE nome ILIKE '%' || $1 || '%';

-- name: FindProdutoByGenero :many
SELECT * FROM t_produtos WHERE genero ILIKE '%' || $1 || '%';

-- name: CreateProduto :one
INSERT INTO t_produtos (nome, descricao, genero)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProduto :one
UPDATE t_produtos 
SET nome = $1, descricao = $2, genero = $3
WHERE id = $4
RETURNING *;

-- name: DeleteProdutoById :execrows
DELETE FROM t_produtos WHERE id = $1 
RETURNING id;
