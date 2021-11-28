package service

import (
	"context"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/domain/model"
)

//thingService struct
type thingService struct {
	r domain.Repository
}

//NewThingService builder
func NewThingService(dr domain.Repository) domain.Service {
	return &thingService{
		r: dr,
	}
}

//Get service implemetation
func (s *thingService) Get(ctx context.Context, id int64) ([]*model.Thing, error) {

	rr, err := s.r.Retrieve(ctx, id)

	if err != nil {
		return nil, err
	}

	return rr, nil
}

//Update service
func (s *thingService) Update(ctx context.Context, t *model.Thing) (*model.Thing, error) {

	r, err := s.r.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	return r, nil
}

//Insert service
func (s *thingService) Insert(ctx context.Context, t *model.Thing) (*model.Thing, error) {

	r, err := s.r.Create(ctx, t)

	if err != nil {
		return nil, err
	}

	return r, nil
}

//Delete service
func (s *thingService) Delete(ctx context.Context, id int64) (*model.Thing, error) {

	r, err := s.r.Delete(ctx, id)

	if err != nil {
		return nil, err
	}

	return r, nil
}
