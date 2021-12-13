package api

import (
	"fmt"
	"net/http"
	"proyectos/src/api/errors"
	"proyectos/src/api/project/api/dto"
	"proyectos/src/api/project/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type ProjectHandler struct {
	domain.Service
}

// GetAll ProjectGetter godoc
// @Summary      Get all projects
// @Description  Get all the projects in the system
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project body dto.Project true "Get all projects"
// @Success      200  {object}  dto.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects [get]
func (ph *ProjectHandler) GetAll(g *gin.Context) {

	dm, err := ph.Service.GetAll(g)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.MapToProjects(dm))
}

// GetByID ProjectGetterByID godoc
// @Summary      Get a project
// @Description  Get a specific project based on it's ID
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project body dto.Project true "Get a project"
// @Success      200  {object}  dto.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects/:id [get]
func (ph *ProjectHandler) GetByID(g *gin.Context) {

	dm, err := ph.Service.GetById(g, g.Param("id"))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err))
		return
	}

	g.JSON(http.StatusOK, dto.FromModel(dm))
}

// Post ProjectCreator godoc
// @Summary      Add a project
// @Description  Add a project to the system
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project body dto.Project true "Create a project"
// @Success      200  {object}  dto.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects [post]
func (ph *ProjectHandler) Post(g *gin.Context) {

	dp := dto.Project{}
	err := g.BindJSON(&dp)
	if err != nil {
		return
	}

	validate := validator.New()
	valerr := validate.StructExcept(dp, "ID")

	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valerr))
		return
	}

	dm, valerr := ph.Service.Insert(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valerr))
		return
	}
	g.JSON(http.StatusOK, dto.FromModel(dm))

}

// Patch ProjectStateUpdater godoc
// @Summary      Update a project's state
// @Description  Update a project's state that is already in the system
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project body dto.Project true "Update a project's state"
// @Success      200  {object}  dto.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects/:id/state [patch]
func (ph *ProjectHandler) Patch(g *gin.Context) {

	dp := dto.Project{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot parse ID",
		})
		return
	}
	dp.ID = i
	g.BindJSON(&dp)

	validate := validator.New()

	valerr := validate.StructPartial(dp, "state")

	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valerr))
		return
	}

	dm, err := ph.Service.Update(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot Save",
		})
		return
	}
	g.JSON(http.StatusOK, dto.FromModel(dm))
}

// Put ProjectPatcher godoc
// @Summary      Modify a project
// @Description  Modify a project that is already in the system
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project body dto.Project true "Create a project"
// @Success      200  {object}  dto.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects/:id [put]
func (ph *ProjectHandler) Put(g *gin.Context) {

	dp := dto.Project{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	dp.ID = i
	g.BindJSON(&dp)

	validate := validator.New()

	valerr := validate.Struct(dp)
	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valerr))
		return
	}

	dm, err := ph.Service.Update(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot Save",
		})
		return
	}
	g.JSON(http.StatusOK, dto.FromModel(dm))
}

// Delete ProjectDeleter godoc
// @Summary      Delete a project
// @Description  Delete a project that is already on the system
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project body dto.Project true "Update a task"
// @Success      200  {object}  dto.Project
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects/:id [delete]
func (ph *ProjectHandler) Delete(g *gin.Context) {
	dp := dto.Project{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.ErrResponse{
			Err:     err,
			Message: "Cannot parse ID",
		})
		return
	}

	dp.ID = i

	_, err = ph.Service.Delete(g, dp.ToModel())

	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrResponse{
			Err:     err,
			Message: "Cannot delete project",
		})
		return
	}

	g.JSON(http.StatusOK, map[string]string{"code": strconv.FormatInt(http.StatusOK, 10), "message": "Project " + g.Param("id") + " deleted successfully"})

}
