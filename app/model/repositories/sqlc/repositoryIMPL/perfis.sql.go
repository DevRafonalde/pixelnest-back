// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: perfis.sql

package repositoryIMPL

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPerfil = `-- name: CreatePerfil :one
INSERT INTO t_perfis (nome, descricao, ativo, data_ultima_atualizacao)
VALUES ($1, $2, $3, $4)
RETURNING id, nome, descricao, ativo, data_ultima_atualizacao
`

type CreatePerfilParams struct {
	Nome                  string
	Descricao             string
	Ativo                 pgtype.Bool
	DataUltimaAtualizacao pgtype.Timestamp
}

func (q *Queries) CreatePerfil(ctx context.Context, arg CreatePerfilParams) (TPerfi, error) {
	row := q.db.QueryRow(ctx, createPerfil,
		arg.Nome,
		arg.Descricao,
		arg.Ativo,
		arg.DataUltimaAtualizacao,
	)
	var i TPerfi
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Ativo,
		&i.DataUltimaAtualizacao,
	)
	return i, err
}

const findAllPerfis = `-- name: FindAllPerfis :many
SELECT id, nome, descricao, ativo, data_ultima_atualizacao FROM t_perfis
`

func (q *Queries) FindAllPerfis(ctx context.Context) ([]TPerfi, error) {
	rows, err := q.db.Query(ctx, findAllPerfis)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TPerfi
	for rows.Next() {
		var i TPerfi
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Descricao,
			&i.Ativo,
			&i.DataUltimaAtualizacao,
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

const findPerfilByID = `-- name: FindPerfilByID :one
SELECT id, nome, descricao, ativo, data_ultima_atualizacao FROM t_perfis WHERE id = $1
`

func (q *Queries) FindPerfilByID(ctx context.Context, id int32) (TPerfi, error) {
	row := q.db.QueryRow(ctx, findPerfilByID, id)
	var i TPerfi
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Ativo,
		&i.DataUltimaAtualizacao,
	)
	return i, err
}

const findPerfilByNome = `-- name: FindPerfilByNome :one
SELECT id, nome, descricao, ativo, data_ultima_atualizacao FROM t_perfis WHERE nome = $1
`

func (q *Queries) FindPerfilByNome(ctx context.Context, nome string) (TPerfi, error) {
	row := q.db.QueryRow(ctx, findPerfilByNome, nome)
	var i TPerfi
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Ativo,
		&i.DataUltimaAtualizacao,
	)
	return i, err
}

const findPerfilByPermissao = `-- name: FindPerfilByPermissao :many
SELECT p.id, p.nome, p.descricao, p.ativo, p.data_ultima_atualizacao 
FROM t_perfis p
JOIN t_perfil_permissao pp ON p.id = pp.perfil_id
WHERE pp.permissao_id = $1
`

func (q *Queries) FindPerfilByPermissao(ctx context.Context, permissaoID int32) ([]TPerfi, error) {
	rows, err := q.db.Query(ctx, findPerfilByPermissao, permissaoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TPerfi
	for rows.Next() {
		var i TPerfi
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Descricao,
			&i.Ativo,
			&i.DataUltimaAtualizacao,
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

const updatePerfil = `-- name: UpdatePerfil :one
UPDATE t_perfis 
SET nome = $1, descricao = $2, ativo = $3, data_ultima_atualizacao = $4
WHERE id = $5
RETURNING id, nome, descricao, ativo, data_ultima_atualizacao
`

type UpdatePerfilParams struct {
	Nome                  string
	Descricao             string
	Ativo                 pgtype.Bool
	DataUltimaAtualizacao pgtype.Timestamp
	ID                    int32
}

func (q *Queries) UpdatePerfil(ctx context.Context, arg UpdatePerfilParams) (TPerfi, error) {
	row := q.db.QueryRow(ctx, updatePerfil,
		arg.Nome,
		arg.Descricao,
		arg.Ativo,
		arg.DataUltimaAtualizacao,
		arg.ID,
	)
	var i TPerfi
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Ativo,
		&i.DataUltimaAtualizacao,
	)
	return i, err
}
