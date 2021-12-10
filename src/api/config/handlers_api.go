package config

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	_ "proyectos/src/api/docs"
	"proyectos/src/api/task/api"
	"proyectos/src/api/task/domain"
	"proyectos/src/api/task/repository"
	"proyectos/src/api/task/service"
)

// @title           PSA Projects API
// @version         1.0
// @description     This API gives access to the projects module.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080


//AddHandlers routes
func (r *SRV) AddHandlers(db *gorm.DB) *SRV {
	tr := repository.NewTaskRepository(db)
	ts := service.NewTaskService(tr)
	r = AddTaskHandler(r, ts)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

//AddTaskHandler routes set
func AddTaskHandler(r *SRV, ds domain.Service) *SRV {
	taskHandler := &api.TaskHandler{
		Service: ds,
	}

	r.POST("/tasks", taskHandler.Post)

	return r
}
