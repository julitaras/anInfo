package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Service interface
type Service interface {
	Get(context.Context, int64) ([]*model.Thing, error)
	Update(context.Context, *model.Thing) (*model.Thing, error)
	Insert(context.Context, *model.Thing) (*model.Thing, error)
	Delete(context.Context, int64) (*model.Thing, error)
}
