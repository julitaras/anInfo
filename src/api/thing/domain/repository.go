package domain

import (
	"context"
	"proyectos/src/api/thing/domain/model"
)

//Repository interface
type Repository interface {
	Ping() error
	Create(context.Context, *model.Thing) (*model.Thing, error)
	Retrieve(context.Context, int64) ([]*model.Thing, error)
	Update(context.Context, *model.Thing) (*model.Thing, error)
	Delete(context.Context, int64) (*model.Thing, error)
}
