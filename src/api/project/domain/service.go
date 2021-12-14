package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Service interface
type Service interface {
	GetAll(context.Context) ([]*model.Projects, error)
	GetById(context.Context, string) (*model.Projects, error)
	Insert(context.Context, *model.Projects) (*model.Projects, error)
	Update(context.Context, *model.Projects) (*model.Projects, error)
	Delete(context.Context, *model.Projects) (*model.Projects, error)
}
