package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"proyectos/src/api/config"
	"proyectos/src/api/config/database"
	"proyectos/src/api/config/settings"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v\n", err)
	}

	r := gin.Default()
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"*"}
	configCors.AllowHeaders = []string{"*"}
	r.Use(cors.New(configCors))

	config.NewServer(r).AddHandlers(db).Run(settings.GetData())

}
