package api

import (
	"net/http"
	"proyectos/src/api/errors"
	"proyectos/src/api/task/api/dto"
	"proyectos/src/api/task/domain"
	"strconv"

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

func (dh *TaskHandler) Delete(g *gin.Context) {
	dp := dto.Task{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot parse ID",
		})
		return
	}

	dp.ID = i

	dm, err := dh.Service.Delete(g, dp.ToModel())

	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot delete task",
		})
		return
	}

	g.JSON(http.StatusOK, dto.FromModel(dm))

}
