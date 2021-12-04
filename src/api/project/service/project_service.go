package service

import (
	"context"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/domain/model"
)

type ProjectService struct {
	r domain.Repository
}

func NewProjectService(dr domain.Repository) domain.Service {
	return &ProjectService{
		r: dr,
	}
}

//Insert service
func (s *ProjectService) Insert(ctx context.Context, t *model.Projects) (*model.Projects, error) {

	r, err := s.r.Create(ctx, t)

	if err != nil {
		return nil, err
	}

	return r, nil
}

//Update service
func (s *ProjectService) Update(ctx context.Context, t *model.Projects) (*model.Projects, error) {

	r, err := s.r.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	return r, nil
}
