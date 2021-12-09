package dto

import (
	"proyectos/src/api/errors"
	"proyectos/src/api/project/domain/model"
	"time"
)

type State string

const (
	Done       State = "DONE"
	InProgress       = "IN_PROGRESS"
	ToDo             = "TODO"
)

type Project struct {
	ID          int64     `validate:"gt=0" json:"id"`
	Name        string    `validate:"required,min=2,max=100" json:"name"`
	Description string    `validate:"required" json:"description"`
	StartDate   time.Time `validate:"required" json:"start_date"`
	FinishDate  time.Time `validate:"required" json:"finish_date"`
	WorkedHours int       `json:"worked_hours"`
	Leader      string    `json:"leader"`
	State       string    `validate:"required" json:"state"`
}

func (p *Project) ToModel() *model.Projects {

	return &model.Projects{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		StartDate:   p.StartDate,
		FinishDate:  p.FinishDate,
		WorkedHours: p.WorkedHours,
		Leader:      p.Leader,
		State:       p.State,
	}
}

func FromModel(modelProject *model.Projects) *Project {
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

func (p *Project) ValidateState() error {
	if !State(p.State).IsValid() {
		return errors.NewErrInvalidState(p.State)
	}
	return nil
}

func (s State) IsValid() bool {
	switch s {
	case Done, InProgress, ToDo:
		return true
	}
	return false
}
