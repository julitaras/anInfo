package domain

import (
	"context"
	"proyectos/src/api/task/domain/model"
)

//Repository interface
type Repository interface {
	Create(context.Context, *model.Task) (*model.Task, error)
}
