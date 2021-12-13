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

func (t *TaskRepository) Create(_ context.Context, task *model.Tasks) (*model.Tasks, error) {
	err := t.DB.Create(&task).Error
	if err != nil {
		log.Printf("Error creating Tasks %v", err)
		return nil, err
	}

	return task, nil
}

func (t *TaskRepository) Update(_ context.Context, task *model.Tasks) (*model.Tasks, error) {
	err := t.DB.Model(task).Updates(task).Error
	if err != nil {
		log.Printf("Error updating Task %v", err)
		return nil, err
	}

	err = t.DB.Find(&task).Where("id = ?", task.ID).Error
	if err != nil {
		log.Printf("Error finding Task %v", err)
		return nil, err
	}

	return task, nil
}

func (t *TaskRepository) Delete(_ context.Context, task *model.Tasks) (*model.Tasks, error) {
	err := t.DB.Delete(&task).Error
	if err != nil {
		log.Printf("Error deleting Tasks %v", err)
		return nil, err
	}

	return task, nil
}

func (t *TaskRepository) GetAll(_ context.Context) ([]*model.Tasks, error) {
	var tasks []*model.Tasks

	err := t.DB.Order("id desc").Find(&tasks).Error
	if err != nil {
		log.Printf("Error getting Tasks %v", err)
		return nil, err
	}

	return tasks, nil
}

func (t *TaskRepository) GetById(_ context.Context, id string) (*model.Tasks, error) {
	var task *model.Tasks

	err := t.DB.First(&task, id).Error
	if err != nil {
		log.Printf("Error getting Task %v", err)
		return nil, err
	}

	return task, nil
}
