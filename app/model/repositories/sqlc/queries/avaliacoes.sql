-- name: FindAllAvaliacoes :many
SELECT * FROM t_avaliacoes;

-- name: FindAvaliacaoById :one
SELECT * FROM t_avaliacoes WHERE id = $1;

-- name: FindAvaliacaoByUsuario :many
SELECT * FROM t_avaliacoes WHERE usuario_id = $1;

-- name: FindAvaliacaoByProduto :many
SELECT * FROM t_avaliacoes WHERE produto_id = $1;

-- name: FindAvaliacaoByJogo :many
SELECT * FROM t_avaliacoes WHERE jogo_id = $1;

-- name: CreateCidade :one
INSERT INTO t_avaliacoes (usuario_id, produto_id, jogo_id, nota, avaliacao) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: UpdateCidade :one
UPDATE t_avaliacoes 
SET usuario_id = $1, produto_id = $2, jogo_id = $3, nota = $4, avaliacao = $5
WHERE id = $6
RETURNING *;

-- name: DeleteCidadeById :execrows
DELETE FROM t_avaliacoes WHERE id = $1 
RETURNING id;
