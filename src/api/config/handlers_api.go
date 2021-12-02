package config

import (
	"gorm.io/gorm"
	"proyectos/src/api/task/api"
	"proyectos/src/api/task/domain"
	"proyectos/src/api/task/repository"
	"proyectos/src/api/task/service"
)

//AddHandlers routes
func (r *SRV) AddHandlers(db *gorm.DB) *SRV {
	tr := repository.NewTaskRepository(db)
	ts := service.NewTaskService(tr)

	r = AddTaskHandler(r, ts)
	return r
}

//AddTaskHandler routes set
func AddTaskHandler(r *SRV, ds domain.Service) *SRV {
	thingHandler := &api.TaskHandler{
		Service: ds,
	}

	r.POST("/task", thingHandler.Post)

	return r
}
