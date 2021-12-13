package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Repository interface
type Repository interface {
	GetAll(context.Context) ([]*model.Projects, error)
	GetById(context.Context, string) (*model.Projects, error)
	Create(context.Context, *model.Projects) (*model.Projects, error)
	Update(context.Context, *model.Projects) (*model.Projects, error)
}
