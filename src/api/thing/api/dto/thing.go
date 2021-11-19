package dto

import (
	"proyectos/src/api/thing/domain/model"
	"time"
)

// Thing dto
type Thing struct {
	ID   int64  `validate:"gt=0"`
	Name string `validate:"required,min=2,max=50"`
}

//ToModel Thing to Model
func (t *Thing) ToModel() *model.Thing { ///////////////// VER!

	return &model.Thing{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}
}

//FromModel Model from model
func (t *Thing) FromModel(dm *model.Thing) {
	t.ID = dm.ID
	t.Name = dm.Name
}
