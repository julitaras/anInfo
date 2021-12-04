package model

import (
	"time"

	"gorm.io/gorm"
)

type Projects struct {
	gorm.Model
	ID          int64 `gorm:"primary_key"`
	Name        string
	Description string
	StartDate   time.Time
	FinishDate  time.Time
	WorkedHours int
	Leader      string
	State       string
}
