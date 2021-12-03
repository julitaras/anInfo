package model

import (
	"time"
)

type Project struct {
	Code        int
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	WorkedHours int
}
