package api

import (
	"fmt"
	"net/http"
	_ "proyectos/src/api/docs"
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

// Post TaskCreator godoc
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

// Put TaskUpdater godoc
// @Summary      Update a task
// @Description  Update a task that is already on the system
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.Task true "Update a task"
// @Success      200  {object}  dto.Task
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks/:id [put]
func (dh *TaskHandler) Put(g *gin.Context) {

	dp := dto.Task{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	dp.ID = i

	err = g.BindJSON(&dp)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err))
		return
	}

	dm, err := dh.Service.Update(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}
	g.JSON(http.StatusOK, dto.FromModel(dm))
}

// Delete TaskDeleter godoc
// @Summary      Delete a task
// @Description  Delete a task that is already on the system
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.Task true "Update a task"
// @Success      200  {object}  dto.Task
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks/:id [delete]
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

// GetAll TaskGetter godoc
// @Summary      Get all tasks
// @Description  Gat all the tasks
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.Task true "Get all tasks"
// @Success      200  {object}  dto.Task
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks [get]
func (dh *TaskHandler) GetAll(g *gin.Context) {

	dm, err := dh.Service.GetAll(g)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.MapToTasks(dm))
}

// GetByID TaskGetterByID godoc
// @Summary      Get a task
// @Description  Gat a specific task based on it's ID
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.Task true "Get a task"
// @Success      200  {object}  dto.Task
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks/:id [get]
func (dh *TaskHandler) GetByID(g *gin.Context) {

	dm, err := dh.Service.GetById(g, g.Param("id"))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.FromModel(dm))
}
