package service

import (
	"simfonia-golang-back/app/model"

	"github.com/devfeel/mapper"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type NumeroTelefonicoService struct {
	db *gorm.DB
}

func NewNumeroTelefonicoService(db *gorm.DB) *NumeroTelefonicoService {
	return &NumeroTelefonicoService{
		db: db,
	}
}

func (NumeroTelefonicoService NumeroTelefonicoService) FindNumeroTelefonicoById(id uint64) (model.NumeroTelefonico, error) {
	numeroTelefonico := new(model.NumeroTelefonico)
	resp := NumeroTelefonicoService.db.First(&numeroTelefonico, id)

	if resp.Error != nil {
		return model.NumeroTelefonico{}, resp.Error
	}

	return *numeroTelefonico, nil
}

func (NumeroTelefonicoService NumeroTelefonicoService) FindNumeroTelefonicoByNumero(numero int) (model.NumeroTelefonico, error) {
	numeroTelefonico := new(model.NumeroTelefonico)
	resp := NumeroTelefonicoService.db.Where("numero = ?", numero).First(&numeroTelefonico)

	if resp.Error != nil {
		return model.NumeroTelefonico{}, resp.Error
	}

	return *numeroTelefonico, nil
}

func (NumeroTelefonicoService NumeroTelefonicoService) FindNumeroTelefonicoBySimCardId(id int) (model.NumeroTelefonico, error) {
	numeroTelefonico := new(model.NumeroTelefonico)
	resp := NumeroTelefonicoService.db.Where("simcad_id = ?", id).First(&numeroTelefonico)

	if resp.Error != nil {
		return model.NumeroTelefonico{}, resp.Error
	}

	return *numeroTelefonico, nil
}

// func (NumeroTelefonicoService NumeroTelefonicoService) FindNumeroTelefonicoBySimCard(simCard model.SimCard) (model.NumeroTelefonico, error) {
// 	numeroTelefonico := new(model.NumeroTelefonico)
// 	resp := NumeroTelefonicoService.db.Where("simcad_id = ?", simCard.ID).First(&numeroTelefonico)

// 	if resp.Error != nil {
// 		return model.NumeroTelefonico{}, resp.Error
// 	}

// 	return *numeroTelefonico, nil
// }

func (NumeroTelefonicoService NumeroTelefonicoService) FindAllNumeroTelefonicos() ([]model.NumeroTelefonico, error) {
	numeroTelefonicos := []model.NumeroTelefonico{}
	resp := NumeroTelefonicoService.db.Find(&numeroTelefonicos)

	if resp.Error != nil {
		return []model.NumeroTelefonico{}, resp.Error
	}

	return numeroTelefonicos, nil
}

func (NumeroTelefonicoService NumeroTelefonicoService) CreateNumeroTelefonico(numeroTelefonico model.NumeroTelefonico) (model.NumeroTelefonico, error) {
	errValidate := validateNumeroTelefonico(numeroTelefonico)
	if errValidate != nil {
		return model.NumeroTelefonico{}, errValidate
	}

	resp := NumeroTelefonicoService.db.Create(&numeroTelefonico)
	if resp.Error != nil {
		return model.NumeroTelefonico{}, resp.Error
	}

	return numeroTelefonico, nil
}

func (NumeroTelefonicoService NumeroTelefonicoService) UpdateNumeroTelefonico(numeroTelefonicoRecebido model.NumeroTelefonico, id uint64) (model.NumeroTelefonico, error) {
	numeroTelefonicoBanco := new(model.NumeroTelefonico)
	respBusca := NumeroTelefonicoService.db.First(&numeroTelefonicoBanco, id)

	if respBusca.Error != nil {
		return model.NumeroTelefonico{}, respBusca.Error
	}

	errValidate := validateNumeroTelefonico(numeroTelefonicoRecebido)
	if errValidate != nil {
		return model.NumeroTelefonico{}, errValidate
	}

	numeroTelefonicoRecebido.ID = numeroTelefonicoBanco.ID
	err := mapper.AutoMapper(numeroTelefonicoRecebido, numeroTelefonicoBanco)
	if err != nil {
		return model.NumeroTelefonico{}, err
	}

	// numeroTelefonicoBanco.CodArea = numeroTelefonicoRecebido.CodArea
	// numeroTelefonicoBanco.Numero = numeroTelefonicoRecebido.Numero
	// numeroTelefonicoBanco.Utilizavel = numeroTelefonicoRecebido.Utilizavel
	// numeroTelefonicoBanco.PortadoIn = numeroTelefonicoRecebido.PortadoIn
	// numeroTelefonicoBanco.PortadoInOperadora = numeroTelefonicoRecebido.PortadoInOperadora
	// numeroTelefonicoBanco.PortadoInDate = numeroTelefonicoRecebido.PortadoInDate
	// numeroTelefonicoBanco.CodigoCNL = numeroTelefonicoRecebido.CodigoCNL
	// numeroTelefonicoBanco.CongeladoAte = numeroTelefonicoRecebido.CongeladoAte
	// numeroTelefonicoBanco.ExternalID = numeroTelefonicoRecebido.ExternalID
	// numeroTelefonicoBanco.PortadoOut = numeroTelefonicoRecebido.PortadoOut
	// numeroTelefonicoBanco.PortadoOutOperadora = numeroTelefonicoRecebido.PortadoOutOperadora
	// numeroTelefonicoBanco.PortadoOutDate = numeroTelefonicoRecebido.PortadoOutDate
	// numeroTelefonicoBanco.DataCriacao = numeroTelefonicoRecebido.DataCriacao
	// numeroTelefonicoBanco.SimCardID = numeroTelefonicoRecebido.SimCardID
	// numeroTelefonicoBanco.SimCard = numeroTelefonicoRecebido.SimCard
	// numeroTelefonicoBanco.PortadoInOperadoraID = numeroTelefonicoRecebido.PortadoInOperadoraID
	// numeroTelefonicoBanco.PortadoInOperadoraObj = numeroTelefonicoRecebido.PortadoInOperadoraObj
	// numeroTelefonicoBanco.PortadoOutOperadoraID = numeroTelefonicoRecebido.PortadoOutOperadoraID
	// numeroTelefonicoBanco.PortadoOutOperadoraObj = numeroTelefonicoRecebido.PortadoOutOperadoraObj
	respSalvamento := NumeroTelefonicoService.db.Save(&numeroTelefonicoBanco)

	if respSalvamento.Error != nil {
		return model.NumeroTelefonico{}, respSalvamento.Error
	}
	return *numeroTelefonicoBanco, nil
}

func (NumeroTelefonicoService NumeroTelefonicoService) DeleteNumeroTelefonicoById(id uint64) (bool, error) {
	resp := NumeroTelefonicoService.db.Delete(&model.NumeroTelefonico{}, id)

	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

func validateNumeroTelefonico(numeroTelefonico model.NumeroTelefonico) error {
	validate := validator.New()
	return validate.Struct(numeroTelefonico)
}
