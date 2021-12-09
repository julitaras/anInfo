package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

type ProjectRepository struct {
	DB *gorm.DB
}

//NewProjectRepository builder
func NewProjectRepository(db *gorm.DB) domain.Repository {
	return &ProjectRepository{
		DB: db,
	}
}

//Create project
func (r ProjectRepository) Create(_ context.Context, project *model.Projects) (*model.Projects, error) {
	err := r.DB.Create(&project).Error
	if err != nil {
		log.Printf("Error creating Project %v", err)
		return nil, err
	}

	return project, nil
}
