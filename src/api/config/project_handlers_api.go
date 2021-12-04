package config

import (
	"proyectos/src/api/project/api"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/repository"
	"proyectos/src/api/project/service"

	"gorm.io/gorm"
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
	r.PATCH("/project/finish", projectHandler.Patch)

	return r
}
