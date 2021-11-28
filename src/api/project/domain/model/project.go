package model

import (
	"time"
)

//Thing model
type Thing struct {
	ID        int64
	Name      string
	IsDeleted bool
	UpdatedAt time.Time
	CreatedAt time.Time
}
