package repository

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/domain/model"
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

func (t *ProjectRepository) GetAll(_ context.Context) ([]*model.Projects, error) {
	var projects []*model.Projects

	err := t.DB.Order("id desc").Find(&projects).Error
	if err != nil {
		log.Printf("Error getting Tasks %v", err)
		return nil, err
	}

	return projects, nil
}

func (t *ProjectRepository) GetById(_ context.Context, id string) (*model.Projects, error) {
	var project *model.Projects

	err := t.DB.First(&project, id).Error
	if err != nil {
		log.Printf("Error getting Task %v", err)
		return nil, err
	}

	return project, nil
}

//Create project
func (r *ProjectRepository) Create(_ context.Context, project *model.Projects) (*model.Projects, error) {
	err := r.DB.Create(&project).Error
	if err != nil {
		log.Printf("Error creating Project %v", err)
		return nil, err
	}

	return project, nil
}

//Update project
func (r ProjectRepository) Update(_ context.Context, project *model.Projects) (*model.Projects, error) {
	err := r.DB.Model(project).Updates(project).Error
	if err != nil {
		log.Printf("Error updating Project %v", err)
		return nil, err
	}

	err = r.DB.Find(&project).Where("id = ?", project.ID).Error
	if err != nil {
		log.Printf("Error finding Project %v", err)
		return nil, err
	}

	return project, nil
}