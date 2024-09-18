-- name: FindAllUsuarioPerfis :many
SELECT * FROM t_usuario_perfil;

-- name: FindUsuarioPerfilByID :one
SELECT * FROM t_usuario_perfil WHERE id = $1;

-- name: FindUsuarioPerfilByPerfil :many
SELECT * FROM t_usuario_perfil WHERE perfil_id = $1;

-- name: FindUsuarioPerfilByUsuario :many
SELECT * FROM t_usuario_perfil WHERE usuario_id = $1;

-- name: CreateUsuarioPerfil :one
INSERT INTO t_usuario_perfil (usuario_id, perfil_id, data_hora)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUsuarioPerfil :one
UPDATE t_usuario_perfil 
SET usuario_id = $1, perfil_id = $2, data_hora = $3
WHERE id = $4
RETURNING *;

-- name: DeleteUsuarioPerfilById :execrows
DELETE FROM t_usuario_perfil 
WHERE id = $1 
RETURNING *;
