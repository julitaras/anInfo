package model

import (
	"proyectos/src/api/project/domain/model"
	"time"
)

//Task model
type Task struct {
	ID             int64
	Name           string
	Description    string
	StartDate      time.Time
	HoursWorked    int
	EstimatedHours int
	ProjectID      int64
	State          string
	CreationDate   time.Time
	AssignedTo     string
	Project        model.Project
}
