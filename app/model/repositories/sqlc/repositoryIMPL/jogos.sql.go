// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: jogos.sql

package repositoryIMPL

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createJogo = `-- name: CreateJogo :one
INSERT INTO t_jogos (nome, sinopse, genero)
VALUES ($1, $2, $3)
RETURNING id, nome, sinopse, avaliacao, genero
`

type CreateJogoParams struct {
	Nome    string
	Sinopse pgtype.Text
	Genero  pgtype.Text
}

func (q *Queries) CreateJogo(ctx context.Context, arg CreateJogoParams) (TJogo, error) {
	row := q.db.QueryRow(ctx, createJogo, arg.Nome, arg.Sinopse, arg.Genero)
	var i TJogo
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Sinopse,
		&i.Avaliacao,
		&i.Genero,
	)
	return i, err
}

const deleteJogoById = `-- name: DeleteJogoById :execrows
DELETE FROM t_jogos WHERE id = $1 
RETURNING id
`

func (q *Queries) DeleteJogoById(ctx context.Context, id int32) (int64, error) {
	result, err := q.db.Exec(ctx, deleteJogoById, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllJogos = `-- name: FindAllJogos :many
SELECT id, nome, sinopse, avaliacao, genero FROM t_jogos
`

func (q *Queries) FindAllJogos(ctx context.Context) ([]TJogo, error) {
	rows, err := q.db.Query(ctx, findAllJogos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TJogo
	for rows.Next() {
		var i TJogo
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Sinopse,
			&i.Avaliacao,
			&i.Genero,
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

const findJogoByGenero = `-- name: FindJogoByGenero :many
SELECT id, nome, sinopse, avaliacao, genero FROM t_jogos WHERE genero ILIKE $1
`

func (q *Queries) FindJogoByGenero(ctx context.Context, genero pgtype.Text) ([]TJogo, error) {
	rows, err := q.db.Query(ctx, findJogoByGenero, genero)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TJogo
	for rows.Next() {
		var i TJogo
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Sinopse,
			&i.Avaliacao,
			&i.Genero,
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

const findJogoByID = `-- name: FindJogoByID :one
SELECT id, nome, sinopse, avaliacao, genero FROM t_jogos WHERE id = $1
`

func (q *Queries) FindJogoByID(ctx context.Context, id int32) (TJogo, error) {
	row := q.db.QueryRow(ctx, findJogoByID, id)
	var i TJogo
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Sinopse,
		&i.Avaliacao,
		&i.Genero,
	)
	return i, err
}

const findJogoByNome = `-- name: FindJogoByNome :many
SELECT id, nome, sinopse, avaliacao, genero FROM t_jogos WHERE nome ILIKE $1
`

func (q *Queries) FindJogoByNome(ctx context.Context, nome string) ([]TJogo, error) {
	rows, err := q.db.Query(ctx, findJogoByNome, nome)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TJogo
	for rows.Next() {
		var i TJogo
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Sinopse,
			&i.Avaliacao,
			&i.Genero,
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

const updateJogo = `-- name: UpdateJogo :one
UPDATE t_jogos 
SET nome = $1, sinopse = $2, genero = $3
WHERE id = $4
RETURNING id, nome, sinopse, avaliacao, genero
`

type UpdateJogoParams struct {
	Nome    string
	Sinopse pgtype.Text
	Genero  pgtype.Text
	ID      int32
}

func (q *Queries) UpdateJogo(ctx context.Context, arg UpdateJogoParams) (TJogo, error) {
	row := q.db.QueryRow(ctx, updateJogo,
		arg.Nome,
		arg.Sinopse,
		arg.Genero,
		arg.ID,
	)
	var i TJogo
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Sinopse,
		&i.Avaliacao,
		&i.Genero,
	)
	return i, err
}
