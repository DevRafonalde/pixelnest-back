package service

import (
	"simfonia-golang-back/app/model"

	"github.com/devfeel/mapper"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SimCardEstadoService struct {
	db *gorm.DB
}

func NewSimCardEstadoService(db *gorm.DB) *SimCardEstadoService {
	return &SimCardEstadoService{
		db: db,
	}
}

func (simCardEstadoService SimCardEstadoService) FindSimCardEstadoById(id uint64) (model.SimCardEstado, error) {
	simCardEstado := new(model.SimCardEstado)
	resp := simCardEstadoService.db.First(&simCardEstado, id)

	if resp.Error != nil {
		return model.SimCardEstado{}, resp.Error
	}

	return *simCardEstado, nil
}

func (simCardEstadoService SimCardEstadoService) FindSimCardEstadoByEstado(estado string) (model.SimCardEstado, error) {
	simCardEstado := new(model.SimCardEstado)
	resp := simCardEstadoService.db.Where("estado = ?", estado).First(&simCardEstado)

	if resp.Error != nil {
		return model.SimCardEstado{}, resp.Error
	}

	return *simCardEstado, nil
}

func (simCardEstadoService SimCardEstadoService) FindAllSimCardEstados() ([]model.SimCardEstado, error) {
	simCardEstados := []model.SimCardEstado{}
	resp := simCardEstadoService.db.Find(&simCardEstados)

	if resp.Error != nil {
		return []model.SimCardEstado{}, resp.Error
	}

	return simCardEstados, nil
}

func (simCardEstadoService SimCardEstadoService) CreateSimCardEstado(simCardEstado model.SimCardEstado) (model.SimCardEstado, error) {
	err := validateSimCardEstado(simCardEstado)
	if err != nil {
		return model.SimCardEstado{}, err
	}

	resp := simCardEstadoService.db.Create(&simCardEstado)
	if resp.Error != nil {
		return model.SimCardEstado{}, resp.Error
	}

	return simCardEstado, nil
}

func (simCardEstadoService SimCardEstadoService) UpdateSimCardEstado(simCardEstadoRecebida model.SimCardEstado, id uint64) (model.SimCardEstado, error) {
	simCardEstadoBanco := new(model.SimCardEstado)
	respBusca := simCardEstadoService.db.First(&simCardEstadoBanco, id)

	if respBusca.Error != nil {
		return model.SimCardEstado{}, respBusca.Error
	}

	err := validateSimCardEstado(simCardEstadoRecebida)
	if err != nil {
		return model.SimCardEstado{}, err
	}

	simCardEstadoRecebida.ID = simCardEstadoBanco.ID
	mapper.AutoMapper(simCardEstadoRecebida, simCardEstadoBanco)

	respSalvamento := simCardEstadoService.db.Save(&simCardEstadoBanco)

	if respSalvamento.Error != nil {
		return model.SimCardEstado{}, respSalvamento.Error
	}
	return *simCardEstadoBanco, nil
}

func (simCardEstadoService SimCardEstadoService) DeleteSimCardEstadoById(id uint64) (bool, error) {
	resp := simCardEstadoService.db.Delete(&model.SimCardEstado{}, id)

	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

func validateSimCardEstado(simCard model.SimCardEstado) error {
	validate := validator.New()
	return validate.Struct(simCard)
}
