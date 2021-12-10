package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "proyectos/src/api/docs"
	"proyectos/src/api/errors"
	"proyectos/src/api/task/api/dto"
	"proyectos/src/api/task/domain"
)

//TaskHandler handler
type TaskHandler struct {
	domain.Service
}

// TaskHandler godoc
// @Summary      Add a task
// @Description  Add a task to the system
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.Task true "Create a task"
// @Success      200  {object}  dto.Task
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks [post]
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
