package utils

import "time"

type Task struct {
	ID             int64     `validate:"gt=0" json:"id,omitempty"`
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

type Project struct {
	ID          int64     `validate:"gt=0" json:"id"`
	Name        string    `validate:"required,min=2,max=100" json:"name" example:"Project's name"`
	Description string    `validate:"required" json:"description" example:"Project's description"`
	StartDate   time.Time `validate:"required" json:"start_date" example:"2021-12-14T12:41:09.993-04:00"`
	FinishDate  time.Time `validate:"required" json:"finish_date" example:"2021-12-14T12:41:09.993-04:00"`
	WorkedHours int       `json:"worked_hours"`
	Leader      string    `json:"leader" example:"Project's leader"`
	State       string    `json:"state" example:"Project's state" enums:"TODO,IN_PROGRESS,DONE"`
}
