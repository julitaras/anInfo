package config

import (
	"proyectos/src/api/config/settings"
	"proyectos/src/api/thing/api"
	"proyectos/src/api/thing/domain"
	"proyectos/src/api/thing/repository"
	"proyectos/src/api/thing/service"
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

	group := r.Group("/thing")

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
