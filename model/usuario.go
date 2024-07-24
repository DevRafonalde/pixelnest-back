package model

type Usuario struct {
	ID          uint64 `gorm:"primary_key,autoIncrement"`
	Nome, Email string
}
