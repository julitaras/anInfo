package api

import (
	"fmt"
	"net/http"
	_ "proyectos/src/api/docs"
	"proyectos/src/api/errors"
	"proyectos/src/api/task/api/dto"
	"proyectos/src/api/task/domain"
	"proyectos/src/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

//TaskHandler handler
type TaskHandler struct {
	domain.Service
}

// Post TaskCreator godoc
// @Summary      Create a task
// @Description  Add a task to the system
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.Task true "Create a task"
// @Success      200  {object}  utils.Project
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks [post]
func (dh *TaskHandler) Post(g *gin.Context) {

	dt := dto.Task{}
	err := g.BindJSON(&dt)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Post.bindJson.tasks"))
		return
	}

	validate := validator.New()
	valErr := validate.StructExcept(dt, "ID")

	if valErr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.Post.validate.tasks"))
		return
	}

	if len(dt.State) > 0 {
		err = dt.ValidateState()
		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Post.validateState.tasks"))
			return
		}
	}

	dm, err := dh.Service.Insert(g, dt.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.Post.insert.tasks"))
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
// @Param        id path int true "Task ID"
// @Param        task body dto.Task true "Update a task"
// @Success      200  {object}  utils.Project
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks/:id [put]
func (dh *TaskHandler) Put(g *gin.Context) {

	dt := dto.Task{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	dt.ID = i

	err = g.BindJSON(&dt)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Put.bindJson.tasks"))
		return
	}

	if len(dt.State) > 0 {
		err = dt.ValidateState()
		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Put.validateState.tasks"))
			return
		}
	}

	dm, err := dh.Service.Update(g, dt.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.Put.update.tasks"))
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
// @Param        id path int true "Task ID"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks/:id [delete]
func (dh *TaskHandler) Delete(g *gin.Context) {
	dt := dto.Task{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.Delete.parseID.tasks"))
		return
	}

	dt.ID = i

	_, err = dh.Service.Delete(g, dt.ToModel())

	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Delete.delete.tasks"))
		return
	}

	g.JSON(http.StatusOK, utils.Response{
		Message: "Task " + g.Param("id") + " deleted successfully",
	})
}

// GetAll TaskGetter godoc
// @Summary      Get all tasks
// @Description  Gat all the tasks
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Success      200  {array}  utils.Project
// @Failure      422  {object} errors.ErrResponse
// @Failure      500  {object} errors.ErrResponse
// @Router       /tasks [get]
func (dh *TaskHandler) GetAll(g *gin.Context) {

	dm, err := dh.Service.GetAll(g)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.GetAll.tasks"))
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
// @Param        id path int true "Task ID"
// @Success      200  {object}  utils.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /tasks/:id [get]
func (dh *TaskHandler) GetByID(g *gin.Context) {

	dm, err := dh.Service.GetById(g, g.Param("id"))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.GetByID.tasks"))
		return
	}

	g.JSON(http.StatusOK, dto.FromModel(dm))
}
