package model

import (
	"gorm.io/gorm"
	"time"
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
	Members		string
}

type Member struct {
	gorm.Model
	ID 			int64
}
