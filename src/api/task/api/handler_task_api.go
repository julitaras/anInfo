package api

import (
	"net/http"
	"proyectos/src/api/errors"
	"proyectos/src/api/task/api/dto"
	"proyectos/src/api/task/domain"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

//TaskHandler handler
type TaskHandler struct {
	domain.Service
}

//Post handler
func (dh *TaskHandler) Post(g *gin.Context) {

	dt := dto.Task{}
	err := g.BindJSON(&dt)
	if err != nil {
		return
	}

	validate := validator.New()
	valErr := validate.StructExcept(dt, "ID")

	if valErr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valErr))
		return
	}

	err = dt.ValidateState()
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err))
		return
	}

	dm, err := dh.Service.Insert(g, dt.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.FromModel(dm))
}

//GetAll handler
func (dh *TaskHandler) GetAll(g *gin.Context) {

	dm, err := dh.Service.GetAll(g)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dm)
}

//GetByID handler
func (dh *TaskHandler) GetByID(g *gin.Context) {

	dm, err := dh.Service.GetById(g, "")
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.FromModel(dm))
}
