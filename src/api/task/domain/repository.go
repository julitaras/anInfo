package domain

import (
	"context"
	"proyectos/src/api/task/domain/model"
)

//Repository interface
type Repository interface {
	Create(context.Context, *model.Tasks) (*model.Tasks, error)
	Delete(context.Context, *model.Tasks) (*model.Tasks, error)
}
