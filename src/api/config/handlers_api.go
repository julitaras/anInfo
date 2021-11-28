package config

import (
	"proyectos/src/api/config/settings"
	"proyectos/src/api/project/api"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/repository"
	"proyectos/src/api/project/service"
)

//AddHandlers routes
func (r *SRV) AddHandlers() *SRV {
	//Get configuration
	cfg := settings.GetData()

	//Set repo
	tr := repository.NewThingRepository(cfg)
	//Set service
	ts := service.NewThingService(tr)

	r = AddThingHandler(r, ts)
	return r
}

//AddThingHandler routes set
func AddThingHandler(r *SRV, ds domain.Service) *SRV {

	group := r.Group("/project")

	group.Use() //<-- add middleware

	thingHandler := &api.ThingHandler{
		Service: ds,
	}

	group.GET("/", thingHandler.Get)
	group.POST("/", thingHandler.Post)
	group.PUT("/", thingHandler.Put)
	group.DELETE("/", thingHandler.Delete)

	return r
}
