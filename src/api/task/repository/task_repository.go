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

func (t TaskRepository) Create(_ context.Context, task *model.Tasks) (*model.Tasks, error) {
	err := t.DB.Create(&task).Error
	if err != nil {
		log.Printf("Error creating Tasks %v", err)
		return nil, err
	}

	return task, nil
}

func (r TaskRepository) Update(_ context.Context, task *model.Tasks) (*model.Tasks, error) {
	err := r.DB.Model(task).Updates(task).Error
	if err != nil {
		log.Printf("Error updating Task %v", err)
		return nil, err
	}

	err = r.DB.Find(&task).Where("id = ?", task.ID).Error
	if err != nil {
		log.Printf("Error finding Task %v", err)
		return nil, err
	}

	return task, nil
}