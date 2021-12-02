package domain

import (
	"context"
	"proyectos/src/api/task/domain/model"
)

//Service interface
type Service interface{
	Insert(context.Context, *model.Task) (*model.Task, error)
}
