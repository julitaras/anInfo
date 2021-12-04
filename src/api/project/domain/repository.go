package domain

import (
	"context"
	"proyectos/src/api/project/domain/model"
)

//Repository interface
type Repository interface {
	Create(context.Context, *model.Projects) (*model.Projects, error)
}
