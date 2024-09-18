package model

import (
	"time"
)

type SimCard struct {
	ID                uint64            `gorm:"primaryKey;autoIncrement"`
	ICCID             string            `gorm:"size:20;not null" validate:"required, len=20"`
	IMSI              string            `gorm:"size:15;not null" validate:"required, len=15"`
	PIN               string            `gorm:"size:8;not null" validate:"required, len=8"`
	PUK               string            `gorm:"size:10;not null" validate:"required, len=10"`
	KI                string            `gorm:"size:16;not null" validate:"required, len=16"`
	OPC               string            `gorm:"size:16;not null" validate:"required, len=16"`
	EstadoID          *uint64           `gorm:"size:4" validate:"len=4"`
	Estado            SimCardEstado     `gorm:"foreignKey:EstadoID"`
	TelefoniaNumeroID *uint64           `gorm:"size:8" validate:"len=8"`
	TelefoniaNumero   *NumeroTelefonico `gorm:"foreignKey:TelefoniaNumeroID"`
	DataCriacao       time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	DataEstado        time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	AtualizadoEm      time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	PUK2              string            `gorm:"size:10" validate:"len=10"`
	PIN2              string            `gorm:"size:8" validate:"len=8"`
}

func (SimCard) TableName() string {
	return "t_simcard"
}
