package model

import (
	"gorm.io/gorm"
	"time"
)

type Projects struct {
	gorm.Model
	ID          int64 `gorm:"primary_key"`
	Code        int
	Name        string
	Description string
	StartDate   time.Time
	FinishDate  time.Time
	HoursWorked int
	Leader      string
}
