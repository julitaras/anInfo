package dto

import (
	"proyectos/src/api/errors"
	"proyectos/src/api/task/domain/model"
	"proyectos/src/api/utils"
	"time"
)

// Task dto
type Task struct {
	ID             int64     `validate:"gt=0" json:"id,omitempty" swaggerignore:"true"`
	Name           string    `validate:"required,min=2,max=500" json:"name" example:"task's name"`
	Description    string    `validate:"required" json:"description" example:"task's description"`
	StartDate      time.Time `json:"start_date" example:"2021-12-14T12:41:09.993-04:00"`
	WorkedHours    int       `json:"worked_hours"`
	EstimatedHours int       `json:"estimated_hours"`
	ProjectID      int64     `validate:"required" json:"project_id"`
	State          string    `json:"state" enums:"TODO,IN_PROGRESS,DONE"`
	CreationDate   time.Time `json:"creation_date" example:"2021-12-14T12:41:09.993-04:00"`
	AssignedTo     string    `json:"assigned_to"`
}

//ToModel Task to Model
func (t *Task) ToModel() *model.Tasks {
	state := utils.ToDo
	if len(t.State) > 0 {
		state = t.State
	}

	creationDate := time.Now().UTC()
	if !t.CreationDate.IsZero() {
		creationDate = t.CreationDate
	}

	return &model.Tasks{
		ID:             t.ID,
		Name:           t.Name,
		Description:    t.Description,
		StartDate:      t.StartDate,
		WorkedHours:    t.WorkedHours,
		EstimatedHours: t.EstimatedHours,
		ProjectID:      t.ProjectID,
		State:          state,
		CreationDate:   creationDate,
		AssignedTo:     t.AssignedTo,
	}
}

//FromModel Model from model
func FromModel(dm *model.Tasks) *Task {
	return &Task{
		ID:             dm.ID,
		Name:           dm.Name,
		Description:    dm.Description,
		StartDate:      dm.StartDate,
		WorkedHours:    dm.WorkedHours,
		EstimatedHours: dm.EstimatedHours,
		ProjectID:      dm.ProjectID,
		State:          dm.State,
		CreationDate:   dm.CreationDate,
		AssignedTo:     dm.AssignedTo,
	}
}

//MapToTasks List of tasks to list of dto task
func MapToTasks(dm []*model.Tasks) []*Task {
	var tasks []*Task

	for _, t := range dm {
		tasks = append(tasks, FromModel(t))
	}

	return tasks
}

func (t *Task) ValidateState() error {
	if !utils.State(t.State).IsValid() {
		return errors.NewErrInvalidState(t.State)
	}
	return nil
}
