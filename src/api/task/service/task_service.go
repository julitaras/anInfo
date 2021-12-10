package service

import (
	"context"
	"proyectos/src/api/task/domain"
	"proyectos/src/api/task/domain/model"
)

//TaskService struct
type TaskService struct {
	r domain.Repository
}

//NewTaskService builder
func NewTaskService(dr domain.Repository) domain.Service {
	return &TaskService{
		r: dr,
	}
}

//Insert service
func (s *TaskService) Insert(ctx context.Context, t *model.Tasks) (*model.Tasks, error) {
	r, err := s.r.Create(ctx, t)
	if err != nil {
		return nil, err
	}

	return r, nil
}

//Update service
func (s *TaskService) Update(ctx context.Context, t *model.Tasks) (*model.Tasks, error) {

	r, err := s.r.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	return r, nil
}
