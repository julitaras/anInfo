package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Service interface
type Service interface {
	Select(context.Context, string) (*model.Projects, error)
	Insert(context.Context, *model.Projects) (*model.Projects, error)
	Update(context.Context, *model.Projects) (*model.Projects, error)
}
