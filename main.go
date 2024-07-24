package main

import (
	"crud-rafael/controller"
	"crud-rafael/db"
	"crud-rafael/service"
)

func main() {
	db := db.CreateConnection()
	usuarioService := service.NewUsuarioService(db)
	usuarioController := controller.NewusuarioController(usuarioService)

	usuarioController.InitRoutes()
}
