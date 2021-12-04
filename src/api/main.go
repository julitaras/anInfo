package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"proyectos/src/api/config"
	"proyectos/src/api/config/database"
	"proyectos/src/api/config/settings"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v\n", err)
	}

	srv := config.NewServer(gin.New())
	srv.AddProjectHandlers(db)
	srv.AddTaskHandlers(db).Run(settings.GetData())

}
