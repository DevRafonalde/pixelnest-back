// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: parametros.sql

package repositoryIMPL

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createParametro = `-- name: CreateParametro :one
INSERT INTO t_parametros (nome, valor, descricao) 
VALUES ($1, $2, $3)
RETURNING id, nome, descricao, valor
`

type CreateParametroParams struct {
	Nome      pgtype.Text
	Valor     pgtype.Text
	Descricao pgtype.Text
}

func (q *Queries) CreateParametro(ctx context.Context, arg CreateParametroParams) (TParametro, error) {
	row := q.db.QueryRow(ctx, createParametro, arg.Nome, arg.Valor, arg.Descricao)
	var i TParametro
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Valor,
	)
	return i, err
}

const deleteParametroById = `-- name: DeleteParametroById :execrows
DELETE FROM t_parametros WHERE id = $1 
RETURNING id
`

func (q *Queries) DeleteParametroById(ctx context.Context, id int32) (int64, error) {
	result, err := q.db.Exec(ctx, deleteParametroById, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllParametros = `-- name: FindAllParametros :many
SELECT id, nome, descricao, valor FROM t_parametros
`

func (q *Queries) FindAllParametros(ctx context.Context) ([]TParametro, error) {
	rows, err := q.db.Query(ctx, findAllParametros)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TParametro
	for rows.Next() {
		var i TParametro
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Descricao,
			&i.Valor,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findParametroById = `-- name: FindParametroById :one
SELECT id, nome, descricao, valor FROM t_parametros WHERE id = $1
`

func (q *Queries) FindParametroById(ctx context.Context, id int32) (TParametro, error) {
	row := q.db.QueryRow(ctx, findParametroById, id)
	var i TParametro
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Valor,
	)
	return i, err
}

const findParametroByNome = `-- name: FindParametroByNome :one
SELECT id, nome, descricao, valor FROM t_parametros WHERE nome = $1
`

func (q *Queries) FindParametroByNome(ctx context.Context, nome pgtype.Text) (TParametro, error) {
	row := q.db.QueryRow(ctx, findParametroByNome, nome)
	var i TParametro
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Valor,
	)
	return i, err
}

const updateParametro = `-- name: UpdateParametro :one
UPDATE t_parametros 
SET nome = $1, valor = $2, descricao = $3
WHERE id = $4
RETURNING id, nome, descricao, valor
`

type UpdateParametroParams struct {
	Nome      pgtype.Text
	Valor     pgtype.Text
	Descricao pgtype.Text
	ID        int32
}

func (q *Queries) UpdateParametro(ctx context.Context, arg UpdateParametroParams) (TParametro, error) {
	row := q.db.QueryRow(ctx, updateParametro,
		arg.Nome,
		arg.Valor,
		arg.Descricao,
		arg.ID,
	)
	var i TParametro
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Valor,
	)
	return i, err
}
