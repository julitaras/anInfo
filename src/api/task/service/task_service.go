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

func (s *TaskService) Delete(ctx context.Context, t *model.Tasks) (*model.Tasks, error) {
	r, err := s.r.Delete(ctx, t)
	if err != nil {
		return nil, err
	}

	return r, nil
}

//GetAll service
func (s *TaskService) GetAll(ctx context.Context) ([]*model.Tasks, error) {
	r, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}

//GetById service
func (s *TaskService) GetById(ctx context.Context, id string) (*model.Tasks, error) {
	r, err := s.r.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return r, nil
}
