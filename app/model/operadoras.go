package model

// Operadora representa a estrutura da tabela t_operadoras.
type Operadora struct {
	ID                   uint64             `gorm:"primaryKey;autoIncrement"`
	Nome                 string             `gorm:"unique;not null" validate:"required"`
	Abreviacao           string             `gorm:"unique;not null" validate:"required"`
	NumeroTelefonicosIn  []NumeroTelefonico `gorm:"foreignKey:PortadoInOperadoraID"`
	NumeroTelefonicosOut []NumeroTelefonico `gorm:"foreignKey:PortadoOutOperadoraID"`
}

// TableName define o nome da tabela para GORM
func (Operadora) TableName() string {
	return "t_operadoras"
}
