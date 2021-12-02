package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"proyectos/src/api/task/domain"
	"proyectos/src/api/task/domain/model"
)

//TaskRepository type
type TaskRepository struct {
	DB *gorm.DB
}

//NewTaskRepository builder
func NewTaskRepository(db *gorm.DB) domain.Repository {
	return &TaskRepository{
		DB: db,
	}
}

func (t TaskRepository) Create(_ context.Context, task *model.Task) (*model.Task, error) {
	err := t.DB.Create(&task).Error
	if err != nil {
		log.Fatalf("Error creating Task %v", err)
		return nil, err
	}

	return task, nil
}