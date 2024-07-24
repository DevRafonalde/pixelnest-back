package service

import (
	"crud-rafael/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar-se com o banco de dados")
	}

	db.AutoMigrate(&model.Usuario{})
	return db
}

func TestCreateUsuario(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuario := model.Usuario{Nome: "Usuario Teste", Email: "teste@exemplo.com"}
	usuarioCriado, err := usuarioService.CreateUsuario(usuario)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, usuarioCriado.ID)
	assert.Equal(t, "Usuario Teste", usuarioCriado.Nome)
	assert.Equal(t, "teste@exemplo.com", usuarioCriado.Email)
}

func TestFindUsuarioById(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuario := model.Usuario{Nome: "Usuario Teste", Email: "teste@exemplo.com"}
	usuarioCriado, err := usuarioService.CreateUsuario(usuario)
	assert.Nil(t, err)

	usuarioEncontrado, err := usuarioService.FindUsuarioById(usuarioCriado.ID)
	assert.Nil(t, err)
	assert.Equal(t, "Usuario Teste", usuarioEncontrado.Nome)
	assert.Equal(t, "teste@exemplo.com", usuarioEncontrado.Email)
}

func TestFindUsuarioByEmail(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuario := model.Usuario{Nome: "Usuario Teste", Email: "teste@exemplo.com"}
	_, err := usuarioService.CreateUsuario(usuario)
	assert.Nil(t, err)

	usuarioEncontrado, err := usuarioService.FindUsuarioByEmail("teste@exemplo.com")
	assert.Nil(t, err)
	assert.Equal(t, "Usuario Teste", usuarioEncontrado.Nome)
	assert.Equal(t, "teste@exemplo.com", usuarioEncontrado.Email)
}

func TestFindAllUsuarios(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuarios := []model.Usuario{
		{Nome: "Usuario Teste 1", Email: "teste1@exemplo.com"},
		{Nome: "Usuario Teste 2", Email: "teste2@exemplo.com"},
		{Nome: "Usuario Teste 3", Email: "teste3@exemplo.com"},
	}

	for _, usuario := range usuarios {
		_, err := usuarioService.CreateUsuario(usuario)
		assert.Nil(t, err)
	}

	usuariosEncontrados, err := usuarioService.FindAllUsuarios()
	assert.Nil(t, err)
	assert.Equal(t, len(usuarios), len(usuariosEncontrados))

	for i, usuario := range usuarios {
		assert.Equal(t, usuario.Nome, usuariosEncontrados[i].Nome)
		assert.Equal(t, usuario.Email, usuariosEncontrados[i].Email)
	}
}

func TestUpdateUsuario(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuario := model.Usuario{Nome: "Usuario Teste", Email: "teste@exemplo.com"}
	usuarioCriado, err := usuarioService.CreateUsuario(usuario)
	assert.Nil(t, err)

	usuarioAtualizado := model.Usuario{Nome: "Usuário Atualizado", Email: "atualizado@exemplo.com"}
	atualizado, err := usuarioService.UpdateUsuario(usuarioAtualizado, usuarioCriado.ID)
	assert.Nil(t, err)
	assert.Equal(t, "Usuário Atualizado", atualizado.Nome)
	assert.Equal(t, "atualizado@exemplo.com", atualizado.Email)
}

func TestDeleteUsuarioById(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuario := model.Usuario{Nome: "Usuario Teste", Email: "teste@exemplo.com"}
	usuarioCriado, err := usuarioService.CreateUsuario(usuario)
	assert.Nil(t, err)

	deletado, err := usuarioService.DeleteUsuarioById(usuarioCriado.ID)
	assert.Nil(t, err)
	assert.True(t, deletado)

	_, err = usuarioService.FindUsuarioById(usuarioCriado.ID)
	assert.NotNil(t, err)
}

func TestDeleteAllUsuarios(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuarios := []model.Usuario{
		{Nome: "Usuario Teste 1", Email: "teste1@exemplo.com"},
		{Nome: "Usuario Teste 2", Email: "teste2@exemplo.com"},
		{Nome: "Usuario Teste 3", Email: "teste3@exemplo.com"},
	}

	for _, usuario := range usuarios {
		_, err := usuarioService.CreateUsuario(usuario)
		assert.Nil(t, err)
	}

	err := usuarioService.DeleteAllUsuarios()
	assert.Nil(t, err)

	usuariosEncontrados, err := usuarioService.FindAllUsuarios()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(usuariosEncontrados))
}

func TestCreateUsuarioDuplicateEmail(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuario := model.Usuario{Nome: "Usuario Teste", Email: "teste@exemplo.com"}
	_, err := usuarioService.CreateUsuario(usuario)
	assert.Nil(t, err)

	// Tentar criar o mesmo usuário novamente
	usuarioDuplicado := model.Usuario{Nome: "Usuario Teste 2", Email: "teste@exemplo.com"}
	usuarioCriado, err := usuarioService.CreateUsuario(usuarioDuplicado)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), usuarioCriado.ID)
}

func TestUpdateUsuarioDuplicateEmail(t *testing.T) {
	db := setupTestDB()
	usuarioService := NewUsuarioService(db)

	usuario1 := model.Usuario{Nome: "Usuario Teste 1", Email: "teste1@exemplo.com"}
	usuario2 := model.Usuario{Nome: "Usuario Teste 2", Email: "teste2@exemplo.com"}
	usuarioCriado1, err := usuarioService.CreateUsuario(usuario1)
	assert.Nil(t, err)
	_, err = usuarioService.CreateUsuario(usuario2)
	assert.Nil(t, err)

	// Tentar atualizar usuario1 para ter o mesmo email de usuario2
	usuarioAtualizado := model.Usuario{Nome: "Usuario Teste 1 Atualizado", Email: "teste2@exemplo.com"}
	atualizado, err := usuarioService.UpdateUsuario(usuarioAtualizado, usuarioCriado1.ID)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), atualizado.ID)
}
