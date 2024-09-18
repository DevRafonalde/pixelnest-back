package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=example dbname=simcard_local port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("falha ao se conectar a DB ", err)
	}

	return db
}
