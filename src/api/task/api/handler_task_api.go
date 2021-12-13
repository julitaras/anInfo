package api

import (
	"fmt"
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
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err))
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

func (dh *TaskHandler) Put(g *gin.Context) {

	dp := dto.Task{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	dp.ID = i
	g.BindJSON(&dp)

	dm, err := dh.Service.Update(g, dp.ToModel())
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
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.ErrResponse{
			Err:     err,
			Message: "Cannot parse ID",
		})
		return
	}

	dp.ID = i

	_, err = dh.Service.Delete(g, dp.ToModel())

	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrResponse{
			Err:     err,
			Message: "Cannot delete task",
		})
		return
	}

	g.JSON(http.StatusOK, map[string]string{"code": strconv.FormatInt(http.StatusOK, 10), "message": "Task " + g.Param("id") + " deleted successfully"})

}

//GetAll handler
func (dh *TaskHandler) GetAll(g *gin.Context) {

	dm, err := dh.Service.GetAll(g)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.MapToTasks(dm))
}

//GetByID handler
func (dh *TaskHandler) GetByID(g *gin.Context) {

	dm, err := dh.Service.GetById(g, g.Param("id"))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.FromModel(dm))
}
