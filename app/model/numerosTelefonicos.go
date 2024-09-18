package model

import (
	"time"
)

type NumeroTelefonico struct {
	ID                     uint64     `gorm:"primaryKey;autoIncrement"`
	CodArea                *int       `gorm:"size:4" validate:"len=4"`
	Numero                 uint64     `gorm:"unique;not null" validate:"required"`
	Utilizavel             bool       `gorm:"default:true"`
	PortadoIn              bool       `gorm:"default:false"`
	PortadoInOperadora     *string    `gorm:"size:50" validate:"len=50"`
	PortadoInDate          *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CodigoCNL              string     `gorm:"size:10;not null" validate:"required, len=10"`
	CongeladoAte           *time.Time `gorm:"type:date"`
	ExternalID             *uint64
	PortadoOut             bool       `gorm:"default:false"`
	PortadoOutOperadora    *string    `gorm:"size:50" validate:"len=50"`
	PortadoOutDate         *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DataCriacao            *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	SimCardID              *uint64
	SimCard                *SimCard `gorm:"foreignKey:SimCardID;references:ID"`
	PortadoInOperadoraID   *uint64
	PortadoInOperadoraObj  *Operadora `gorm:"foreignKey:PortadoInOperadoraID;references:ID"`
	PortadoOutOperadoraID  *uint64
	PortadoOutOperadoraObj *Operadora `gorm:"foreignKey:PortadoOutOperadoraID;references:ID"`
}

func (NumeroTelefonico) TableName() string {
	return "t_telefonia_numero"
}
