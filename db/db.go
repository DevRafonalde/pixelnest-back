package db

import (
	"context"
	"crud-rafael/model"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := "postgres://postgres:example@localhost:5432/postgres"
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	createDB(dbpool)
	dsn = "host=localhost user=postgres password=example dbname=usuario port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("falha ao se conectar a DB ", err)
	}

	err = db.AutoMigrate(&model.Usuario{})
	if err != nil {
		log.Fatal("falha ao migrar database ", err)
	}
	return db
}

func createDB(dbpool *pgxpool.Pool) {
	_, err := dbpool.Exec(context.Background(), "CREATE DATABASE usuario")
	if err != nil {
		// Pode ser que o banco já exista, então log apenas uma mensagem
		log.Printf("Database already exists or failed to create: %v\n", err)
	} else {
		log.Println("Database 'usuario' created successfully.")
	}
}
