package service

import (
	"simfonia-golang-back/app/model"

	"github.com/devfeel/mapper"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CidadeService struct {
	db *gorm.DB
}

func NewCidadeService(db *gorm.DB) *CidadeService {
	return &CidadeService{
		db: db,
	}
}

func (cidadeService CidadeService) FindCidadeById(id int) (model.Cidade, error) {
	cidade := new(model.Cidade)
	resp := cidadeService.db.First(&cidade, id)

	if resp.Error != nil {
		return model.Cidade{}, resp.Error
	}

	return *cidade, nil
}

func (cidadeService CidadeService) FindCidadeByNome(nome string) (model.Cidade, error) {
	cidade := new(model.Cidade)
	resp := cidadeService.db.Where("nome = ?", nome).First(&cidade)

	if resp.Error != nil {
		return model.Cidade{}, resp.Error
	}

	return *cidade, nil
}

func (cidadeService CidadeService) FindAllCidades() ([]model.Cidade, error) {
	cidades := []model.Cidade{}
	resp := cidadeService.db.Find(&cidades)

	if resp.Error != nil {
		return []model.Cidade{}, resp.Error
	}

	return cidades, nil
}

func (cidadeService CidadeService) CreateCidade(cidade model.Cidade) (model.Cidade, error) {
	err := validateCidade(cidade)
	if err != nil {
		return model.Cidade{}, err
	}

	resp := cidadeService.db.Create(&cidade)
	if resp.Error != nil {
		return model.Cidade{}, resp.Error
	}

	return cidade, nil
}

func (cidadeService CidadeService) UpdateCidade(cidadeRecebida model.Cidade, id uint64) (model.Cidade, error) {
	cidadeBanco := new(model.Cidade)
	respBusca := cidadeService.db.First(&cidadeBanco, id)

	if respBusca.Error != nil {
		return model.Cidade{}, respBusca.Error
	}

	err := validateCidade(cidadeRecebida)
	if err != nil {
		return model.Cidade{}, err
	}

	cidadeRecebida.ID = cidadeBanco.ID
	mapper.AutoMapper(cidadeRecebida, cidadeBanco)

	// cidadeBanco.UUID = cidadeRecebida.UUID
	// cidadeBanco.Nome = cidadeRecebida.Nome
	// cidadeBanco.CodIBGE = cidadeRecebida.CodIBGE
	// cidadeBanco.UF = cidadeRecebida.UF
	// cidadeBanco.CodArea = cidadeRecebida.CodArea
	respSalvamento := cidadeService.db.Save(&cidadeBanco)

	if respSalvamento.Error != nil {
		return model.Cidade{}, respSalvamento.Error
	}
	return *cidadeBanco, nil
}

func (cidadeService CidadeService) DeleteCidadeById(id uint64) (bool, error) {
	resp := cidadeService.db.Delete(&model.Cidade{}, id)

	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

func validateCidade(simCard model.Cidade) error {
	validate := validator.New()
	return validate.Struct(simCard)
}
