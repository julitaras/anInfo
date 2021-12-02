package model

import "time"

//Task model
type Task struct {
	ID             int64
	Name           string
	Description    string
	StartDate      time.Time
	HoursWorked    int
	EstimatedHours int
	ProjectId      int64
	State          string
	CreationDate   time.Time
	AssignedTo     string
}
