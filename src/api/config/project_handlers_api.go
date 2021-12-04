package config

import (
	"gorm.io/gorm"
	"proyectos/src/api/project/api"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/repository"
	"proyectos/src/api/project/service"
)

//AddProjectHandlers routes
func (r *SRV) AddProjectHandlers(db *gorm.DB) *SRV {
	pr := repository.NewProjectRepository(db)
	ps := service.NewProjectService(pr)

	r = AddProjectHandler(r, ps)
	return r
}

//AddProjectHandler routes set
func AddProjectHandler(r *SRV, ds domain.Service) *SRV {
	projectHandler := &api.ProjectHandler{
		Service: ds,
	}

	r.POST("/project", projectHandler.Post)

	return r
}
