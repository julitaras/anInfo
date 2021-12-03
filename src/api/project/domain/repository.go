package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Repository interface
type Repository interface {
	Ping() error
	Create(context.Context, *model.Project) (*model.Project, error)
	Retrieve(context.Context, int) ([]*model.Project, error)
	Update(context.Context, *model.Project) (*model.Project, error)
	Delete(context.Context, int) (*model.Project, error)
}
