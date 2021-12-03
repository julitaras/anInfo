package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Service interface
type Service interface {
	Get(context.Context, int) ([]*model.Project, error)
	Update(context.Context, *model.Project) (*model.Project, error)
	Insert(context.Context, *model.Project) (*model.Project, error)
	Delete(context.Context, int) (*model.Project, error)
}
