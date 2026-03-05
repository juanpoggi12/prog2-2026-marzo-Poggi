package main

import (
	"log"
	"prog2-2026-marzo-Poggi/database"
	"prog2-2026-marzo-Poggi/handlers"
	"prog2-2026-marzo-Poggi/middleware"
	"prog2-2026-marzo-Poggi/services"

	"github.com/gin-gonic/gin"
)

var (
	handler *handlers.OperacionesHandler
	router  *gin.Engine
)

func main() {
	router = gin.Default()
	dependencies()
	mappingRoutes()

	log.Println("Iniciando el servidor en http://localhost:8080 ...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}

func mappingRoutes() {
	operaciones := router.Group("/operaciones").Use(middleware.ValidarHeader())
	{
		operaciones.POST("", handler.CalcularPromedioVelocidad)
		operaciones.POST("", handler.CalcularProyeccionMejora)
		operaciones.GET("/:rendimiento_actual/:tasa_mejora/:semanas", handler.CalcularVariabilidadRendimiento)

	}
}

func dependencies() {
	var db database.DB = database.NewMongoDB()
	var serv services.OperacionesServicesInterface = services.NewOperacionesService(db)
	handler = handlers.NewOperacionesHandler(serv)
}
