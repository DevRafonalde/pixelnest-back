package model

// SimCardEstado representa a estrutura da tabela t_simcard_estado.
type SimCardEstado struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Estado    string    `gorm:"size:20;not null"`
	Descricao string    `gorm:"type:text"`
	SimCards  []SimCard `gorm:"foreignKey:EstadoID"`
}

// TableName define o nome da tabela para GORM
func (SimCardEstado) TableName() string {
	return "t_simcard_estado"
}
