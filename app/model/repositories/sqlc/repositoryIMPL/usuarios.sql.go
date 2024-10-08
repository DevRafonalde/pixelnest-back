// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: usuarios.sql

package repositoryIMPL

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUsuario = `-- name: CreateUsuario :one
INSERT INTO t_usuarios (nome, email, senha, ativo, token_reset_senha, data_ultima_atualizacao, senha_atualizada)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, nome, email, senha, ativo, token_reset_senha, data_ultima_atualizacao, senha_atualizada
`

type CreateUsuarioParams struct {
	Nome                  string
	Email                 string
	Senha                 string
	Ativo                 pgtype.Bool
	TokenResetSenha       pgtype.Text
	DataUltimaAtualizacao pgtype.Timestamp
	SenhaAtualizada       pgtype.Bool
}

func (q *Queries) CreateUsuario(ctx context.Context, arg CreateUsuarioParams) (TUsuario, error) {
	row := q.db.QueryRow(ctx, createUsuario,
		arg.Nome,
		arg.Email,
		arg.Senha,
		arg.Ativo,
		arg.TokenResetSenha,
		arg.DataUltimaAtualizacao,
		arg.SenhaAtualizada,
	)
	var i TUsuario
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Email,
		&i.Senha,
		&i.Ativo,
		&i.TokenResetSenha,
		&i.DataUltimaAtualizacao,
		&i.SenhaAtualizada,
	)
	return i, err
}

const findAllUsuarios = `-- name: FindAllUsuarios :many
SELECT id, nome, email, senha, ativo, token_reset_senha, data_ultima_atualizacao, senha_atualizada FROM t_usuarios
`

func (q *Queries) FindAllUsuarios(ctx context.Context) ([]TUsuario, error) {
	rows, err := q.db.Query(ctx, findAllUsuarios)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TUsuario
	for rows.Next() {
		var i TUsuario
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Email,
			&i.Senha,
			&i.Ativo,
			&i.TokenResetSenha,
			&i.DataUltimaAtualizacao,
			&i.SenhaAtualizada,
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

const findUsuarioByEmail = `-- name: FindUsuarioByEmail :one
SELECT id, nome, email, senha, ativo, token_reset_senha, data_ultima_atualizacao, senha_atualizada FROM t_usuarios
WHERE email = $1
`

func (q *Queries) FindUsuarioByEmail(ctx context.Context, email string) (TUsuario, error) {
	row := q.db.QueryRow(ctx, findUsuarioByEmail, email)
	var i TUsuario
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Email,
		&i.Senha,
		&i.Ativo,
		&i.TokenResetSenha,
		&i.DataUltimaAtualizacao,
		&i.SenhaAtualizada,
	)
	return i, err
}

const findUsuarioById = `-- name: FindUsuarioById :one
SELECT id, nome, email, senha, ativo, token_reset_senha, data_ultima_atualizacao, senha_atualizada FROM t_usuarios
WHERE id = $1
`

func (q *Queries) FindUsuarioById(ctx context.Context, id int32) (TUsuario, error) {
	row := q.db.QueryRow(ctx, findUsuarioById, id)
	var i TUsuario
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Email,
		&i.Senha,
		&i.Ativo,
		&i.TokenResetSenha,
		&i.DataUltimaAtualizacao,
		&i.SenhaAtualizada,
	)
	return i, err
}

const updateUsuario = `-- name: UpdateUsuario :one
UPDATE t_usuarios
SET nome = $2, email = $3, senha = $4, ativo = $5, token_reset_senha = $6, data_ultima_atualizacao = $7, senha_atualizada = $8
WHERE id = $1
RETURNING id, nome, email, senha, ativo, token_reset_senha, data_ultima_atualizacao, senha_atualizada
`

type UpdateUsuarioParams struct {
	ID                    int32
	Nome                  string
	Email                 string
	Senha                 string
	Ativo                 pgtype.Bool
	TokenResetSenha       pgtype.Text
	DataUltimaAtualizacao pgtype.Timestamp
	SenhaAtualizada       pgtype.Bool
}

func (q *Queries) UpdateUsuario(ctx context.Context, arg UpdateUsuarioParams) (TUsuario, error) {
	row := q.db.QueryRow(ctx, updateUsuario,
		arg.ID,
		arg.Nome,
		arg.Email,
		arg.Senha,
		arg.Ativo,
		arg.TokenResetSenha,
		arg.DataUltimaAtualizacao,
		arg.SenhaAtualizada,
	)
	var i TUsuario
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Email,
		&i.Senha,
		&i.Ativo,
		&i.TokenResetSenha,
		&i.DataUltimaAtualizacao,
		&i.SenhaAtualizada,
	)
	return i, err
}
