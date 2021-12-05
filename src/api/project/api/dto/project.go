package dto

import (
	"proyectos/src/api/project/domain/model"
	"time"
)

type Project struct {
	ID          int64     `validate:"gt=0" json:"id"`
	Name        string    `validate:"required,min=2,max=100" json:"name"`
	Description string    `validate:"required" json:"description"`
	StartDate   time.Time `validate:"required" json:"start_date"`
	FinishDate  time.Time `validate:"required" json:"finish_date"`
	WorkedHours int       `json:"worked_hours"`
	Leader      string    `json:"leader"`
	State       string    `json:"state"`
}

func (project *Project) ToModel() *model.Projects {

	return &model.Projects{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		StartDate:   project.StartDate,
		FinishDate:  project.FinishDate,
		WorkedHours: project.WorkedHours,
		Leader:      project.Leader,
		State:       project.State,
	}
}

func (project *Project) FromModel(modelProject *model.Projects) *Project {
	return &Project{
		ID:          modelProject.ID,
		Name:        modelProject.Name,
		Description: modelProject.Description,
		StartDate:   modelProject.StartDate,
		FinishDate:  modelProject.FinishDate,
		WorkedHours: modelProject.WorkedHours,
		Leader:      modelProject.Leader,
		State:       modelProject.State,
	}
}
