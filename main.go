package main

import (
	"simfonia-golang-back/app/controller"
	"simfonia-golang-back/app/service"
	"simfonia-golang-back/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.CreateConnection()
	app := gin.Default()

	controller.NewCidadeController(service.NewCidadeService(db)).InitRoutes(app)
	controller.NewNumeroTelefonicoController(service.NewNumeroTelefonicoService(db)).InitRoutes(app)
	controller.NewOperadoraController(service.NewOperadoraService(db)).InitRoutes(app)
	controller.NewSimCardController(service.NewSimCardService(db)).InitRoutes(app)
	controller.NewSimCardEstadoController(service.NewSimCardEstadoService(db)).InitRoutes(app)

	app.Run(":8601")

	// Migração das tabelas e configuração de chaves estrangeiras
	// err = db.AutoMigrate(
	// 	&model.Operadora{},
	// 	&model.Cidade{},
	// 	&model.SimCard{},
	// 	&model.SimCardEstado{},
	// 	&model.NumeroTelefonico{},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
