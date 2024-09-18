package service

import (
	"simfonia-golang-back/app/model"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OperadoraService struct {
	db *gorm.DB
}

func NewOperadoraService(db *gorm.DB) *OperadoraService {
	return &OperadoraService{
		db: db,
	}
}

func (operadoraService OperadoraService) FindOperadoraById(id uint64) (model.Operadora, error) {
	operadora := new(model.Operadora)
	resp := operadoraService.db.First(&operadora, id)

	if resp.Error != nil {
		return model.Operadora{}, resp.Error
	}

	return *operadora, nil
}

func (operadoraService OperadoraService) FindOperadoraByNome(nome string) (model.Operadora, error) {
	operadora := new(model.Operadora)
	resp := operadoraService.db.Where("nome = ?", nome).First(&operadora)

	if resp.Error != nil {
		return model.Operadora{}, resp.Error
	}

	return *operadora, nil
}

func (operadoraService OperadoraService) FindOperadoraByAbreviacao(abreviacao string) (model.Operadora, error) {
	operadora := new(model.Operadora)
	resp := operadoraService.db.Where("abreviacao = ?", abreviacao).First(&operadora)

	if resp.Error != nil {
		return model.Operadora{}, resp.Error
	}

	return *operadora, nil
}

func (operadoraService OperadoraService) FindAllOperadoras() ([]model.Operadora, error) {
	operadoras := []model.Operadora{}
	resp := operadoraService.db.Find(&operadoras)

	if resp.Error != nil {
		return []model.Operadora{}, resp.Error
	}

	return operadoras, nil
}

func (operadoraService OperadoraService) CreateOperadora(operadora model.Operadora) (model.Operadora, error) {
	err := validateOperadora(operadora)
	if err != nil {
		return model.Operadora{}, err
	}

	resp := operadoraService.db.Create(&operadora)
	if resp.Error != nil {
		return model.Operadora{}, resp.Error
	}

	return operadora, nil
}

func (operadoraService OperadoraService) UpdateOperadora(operadoraRecebida model.Operadora, id uint64) (model.Operadora, error) {
	operadoraBanco := new(model.Operadora)
	respBusca := operadoraService.db.First(&operadoraBanco, id)

	if respBusca.Error != nil {
		return model.Operadora{}, respBusca.Error
	}

	err := validateOperadora(operadoraRecebida)
	if err != nil {
		return model.Operadora{}, err
	}

	operadoraBanco.Nome = operadoraRecebida.Nome
	operadoraBanco.Abreviacao = operadoraRecebida.Abreviacao
	respSalvamento := operadoraService.db.Save(&operadoraBanco)

	if respSalvamento.Error != nil {
		return model.Operadora{}, respSalvamento.Error
	}
	return *operadoraBanco, nil
}

func (operadoraService OperadoraService) DeleteOperadoraById(id uint64) (bool, error) {
	resp := operadoraService.db.Delete(&model.Operadora{}, id)

	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

func validateOperadora(operadora model.Operadora) error {
	validate := validator.New()
	return validate.Struct(operadora)
}
