package service

import (
	"crud-rafael/model"

	"gorm.io/gorm"
)

type UsuarioService struct {
	db *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{
		db: db,
	}
}

func (usuarioService *UsuarioService) FindUsuarioById(id uint64) (model.Usuario, error) {
	usuario := new(model.Usuario)
	resp := usuarioService.db.First(&usuario, id)

	if resp.Error != nil {
		return model.Usuario{}, resp.Error
	}

	return *usuario, nil
}

func (usuarioService *UsuarioService) FindUsuarioByEmail(email string) (model.Usuario, error) {
	usuario := new(model.Usuario)
	resp := usuarioService.db.Where("email = ?", email).First(&usuario)

	if resp.Error != nil {
		return model.Usuario{}, resp.Error
	}

	return *usuario, nil
}

func (usuarioService *UsuarioService) FindAllUsuarios() ([]model.Usuario, error) {
	usuarios := []model.Usuario{}
	resp := usuarioService.db.Find(&usuarios)

	if resp.Error != nil {
		return []model.Usuario{}, resp.Error
	}

	return usuarios, nil
}

func (usuarioService *UsuarioService) CreateUsuario(usuario model.Usuario) (model.Usuario, error) {
	if usuarioService.verificarSeEmailEmUso(usuario.Email) {
		return model.Usuario{}, nil
	}

	resp := usuarioService.db.Create(&usuario)

	if resp.Error != nil {
		return model.Usuario{}, resp.Error
	}

	return usuario, nil
}

func (usuarioService *UsuarioService) UpdateUsuario(usuarioRecebido model.Usuario, id uint64) (model.Usuario, error) {
	usuarioBanco := new(model.Usuario)
	respBusca := usuarioService.db.First(&usuarioBanco, id)

	if respBusca.Error != nil {
		return model.Usuario{}, respBusca.Error
	}

	if usuarioService.verificarSeEmailEmUso(usuarioRecebido.Email) {
		return model.Usuario{}, nil
	}

	usuarioBanco.Nome = usuarioRecebido.Nome
	usuarioBanco.Email = usuarioRecebido.Email
	respSalva := usuarioService.db.Save(&usuarioBanco)
	if respSalva.Error != nil {
		return model.Usuario{}, respSalva.Error
	}

	return *usuarioBanco, nil
}

func (usuarioService *UsuarioService) DeleteUsuarioById(id uint64) (bool, error) {
	resp := usuarioService.db.Delete(&model.Usuario{}, id)
	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

func (usuarioService *UsuarioService) DeleteAllUsuarios() error {
	resp := usuarioService.db.Exec("DELETE FROM usuarios")
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (usuarioService *UsuarioService) verificarSeEmailEmUso(email string) bool {
	possivelUsuarioExistente, _ := usuarioService.FindUsuarioByEmail(email)

	if (possivelUsuarioExistente != model.Usuario{}) {
		return true
	} else {
		return false
	}
}
