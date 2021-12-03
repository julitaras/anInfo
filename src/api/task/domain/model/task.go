package model

import (
	"gorm.io/gorm"
	"time"
)

//Tasks model
type Tasks struct {
	gorm.Model
	ID             int64 `gorm:"primary_key"`
	Name           string
	Description    string
	StartDate      time.Time
	HoursWorked    int
	EstimatedHours int
	ProjectID      int64  `gorm:"foreignKey:ProjectsID"`
	State          string //TODO: Vale la pena hacer un enum?
	CreationDate   time.Time
	AssignedTo     string
}
