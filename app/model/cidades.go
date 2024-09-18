package model

import (
	"github.com/google/uuid"
)

// Cidade representa a estrutura da tabela t_cidades.
type Cidade struct {
	ID      uint64    `gorm:"primaryKey;autoIncrement"`
	UUID    uuid.UUID `gorm:"type:uuid;"`
	Nome    string    `gorm:"not null" validate:"required"`
	CodIBGE int       `gorm:"unique;not null" validate:"required"`
	UF      string    `gorm:"size:2;not null" validate:"required, len=2"`
	CodArea int       `gorm:"not null" validate:"required"`
}

// TableName define o nome da tabela para GORM
func (Cidade) TableName() string {
	return "t_cidades"
}
