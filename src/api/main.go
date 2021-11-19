package main

import (
	"github.com/gin-gonic/gin"
	"proyectos/src/api/config"
	"proyectos/src/api/config/settings"
)

func main() {
	config.NewServer(gin.New()).AddHandlers().Run(settings.GetData())
}
