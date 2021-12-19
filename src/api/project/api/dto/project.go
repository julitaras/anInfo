package dto

import (
	"proyectos/src/api/errors"
	"proyectos/src/api/project/domain/model"
	"proyectos/src/api/utils"
	"time"
)

type State string

const (
	Done       State = "DONE"
	InProgress       = "IN_PROGRESS"
	ToDo             = "TODO"
)

type Project struct {
	ID          int64     `validate:"gt=0" json:"id" swaggerignore:"true"`
	Name        string    `validate:"required,min=2,max=100" json:"name" example:"Project's name"`
	Description string    `validate:"required" json:"description" example:"Project's description"`
	StartDate   time.Time `validate:"required" json:"start_date" example:"2021-12-14T12:41:09.993-04:00"`
	FinishDate  time.Time `validate:"required" json:"finish_date" example:"2021-12-14T12:41:09.993-04:00"`
	WorkedHours int       `json:"worked_hours"`
	Leader      string    `json:"leader" example:"Project's leader"`
	State       string    `json:"state" example:"Project's state" enums:"TODO,IN_PROGRESS,DONE"`
	Members     []string  `json:"members" example: Project's members'`
}

func (p *Project) ToModel() *model.Projects {
	state := utils.ToDo
	if len(p.State) > 0 {
		state = p.State
	}

	return &model.Projects{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		StartDate:   p.StartDate,
		FinishDate:  p.FinishDate,
		WorkedHours: p.WorkedHours,
		Leader:      p.Leader,
		State:       state,
		Members:	 p.Members,
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
		Members:	 modelProject.Members,
	}
}

func MapToProjects(modelProjects []*model.Projects) []*Project {
	var projects []*Project

	for _, t := range modelProjects {
		projects = append(projects, FromModel(t))
	}

	return projects
}

func (p *Project) ValidateState() error {
	if !utils.State(p.State).IsValid() {
		return errors.NewErrInvalidState(p.State)
	}
	return nil
}
