package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Service interface
type Service interface {
	Insert(context.Context, *model.Projects) (*model.Projects, error)
	Update(context.Context, *model.Projects) (*model.Projects, error)
}
