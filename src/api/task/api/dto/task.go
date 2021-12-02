package dto

import (
	"proyectos/src/api/task/domain/model"
	"time"
)

// Task dto
type Task struct {
	ID             int64     `validate:"gt=0" json:"id"`
	Name           string    `validate:"required,min=2,max=50" json:"name"`
	Description    string    `validate:"required" json:"description"`
	StartDate      time.Time `validate:"required" json:"start_date"`
	HoursWorked    int       `validate:"required" json:"hours_worked"`
	EstimatedHours int       `validate:"required" json:"estimated_hours"`
	ProjectId      int64     `validate:"required" json:"project_id"`
	State          string    `validate:"required" json:"state"`
	CreationDate   time.Time `validate:"required" json:"creation_date"`
	AssignedTo     string    `validate:"required" json:"assigned_to"`
}

//ToModel Task to Model
func (t *Task) ToModel() *model.Task {

	return &model.Task{
		ID:             t.ID,
		Name:           t.Name,
		Description:    t.Description,
		StartDate:      t.StartDate,
		HoursWorked:    t.HoursWorked,
		EstimatedHours: t.EstimatedHours,
		ProjectId:      t.ProjectId,
		State:          t.State,
		CreationDate:   t.CreationDate, //TODO o ver si es la fecha de hoy en dia
		AssignedTo:     t.AssignedTo,
	}
}

//FromModel Model from model
func (t *Task) FromModel(dm *model.Task) {
	t.ID = dm.ID
	t.Name = dm.Name
	t.Description = dm.Description
	t.StartDate = dm.StartDate
	t.HoursWorked = dm.HoursWorked
	t.EstimatedHours = dm.EstimatedHours
	t.ProjectId = dm.ProjectId
	t.State = dm.State
	t.CreationDate = dm.CreationDate
	t.AssignedTo = dm.AssignedTo
}
