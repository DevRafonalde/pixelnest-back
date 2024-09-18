package service

import (
	"simfonia-golang-back/app/model"

	"github.com/devfeel/mapper"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SimCardService struct {
	db *gorm.DB
}

func NewSimCardService(db *gorm.DB) *SimCardService {
	return &SimCardService{
		db: db,
	}
}

func (simCardService SimCardService) FindSimCardById(id uint64) (model.SimCard, error) {
	simCard := new(model.SimCard)
	resp := simCardService.db.First(&simCard, id)

	if resp.Error != nil {
		return model.SimCard{}, resp.Error
	}

	return *simCard, nil
}

func (SimCardService SimCardService) FindSimCardByTelefoniaNumeroID(id int) (model.SimCard, error) {
	simCard := new(model.SimCard)
	resp := SimCardService.db.Where("simcad_id = ?", id).First(&simCard)

	if resp.Error != nil {
		return model.SimCard{}, resp.Error
	}

	return *simCard, nil
}

// func (SimCardService SimCardService) FindSimCardByTelefoniaNumero(numeroTelefonico model.NumeroTelefonico) (model.SimCard, error) {
// 	simCard := new(model.SimCard)
// 	resp := SimCardService.db.Where("simcad_id = ?", numeroTelefonico.ID).First(&simCard)

// 	if resp.Error != nil {
// 		return model.SimCard{}, resp.Error
// 	}

// 	return *simCard, nil
// }

func (simCardService SimCardService) FindAllSimCards() ([]model.SimCard, error) {
	simCards := []model.SimCard{}
	resp := simCardService.db.Find(&simCards)

	if resp.Error != nil {
		return []model.SimCard{}, resp.Error
	}

	return simCards, nil
}

func (simCardService SimCardService) CreateSimCard(simCard model.SimCard) (model.SimCard, error) {
	err := validateSimCard(simCard)
	if err != nil {
		return model.SimCard{}, err
	}

	resp := simCardService.db.Create(&simCard)
	if resp.Error != nil {
		return model.SimCard{}, resp.Error
	}

	return simCard, nil
}

func (simCardService SimCardService) UpdateSimCard(simCardRecebida model.SimCard, id uint64) (model.SimCard, error) {
	simCardBanco := new(model.SimCard)
	respBusca := simCardService.db.First(&simCardBanco, id)

	if respBusca.Error != nil {
		return model.SimCard{}, respBusca.Error
	}

	err := validateSimCard(simCardRecebida)
	if err != nil {
		return model.SimCard{}, err
	}

	simCardRecebida.ID = simCardBanco.ID
	mapper.AutoMapper(simCardRecebida, simCardBanco)

	respSalvamento := simCardService.db.Save(&simCardBanco)

	if respSalvamento.Error != nil {
		return model.SimCard{}, respSalvamento.Error
	}

	return *simCardBanco, nil
}

func (simCardService SimCardService) DeleteSimCardById(id uint64) (bool, error) {
	resp := simCardService.db.Delete(&model.SimCard{}, id)

	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

func validateSimCard(simCard model.SimCard) error {
	validate := validator.New()
	return validate.Struct(simCard)
}
