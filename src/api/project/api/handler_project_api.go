package api

import (
	"fmt"
	"net/http"
	"proyectos/src/api/errors"
	"proyectos/src/api/project/api/dto"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/utils"
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
// @Success      200  {array}  utils.Project
// @Failure      422  {object} errors.ErrResponse
// @Failure      500  {object} errors.ErrResponse
// @Router       /projects [get]
func (ph *ProjectHandler) GetAll(g *gin.Context) {

	dm, err := ph.Service.GetAll(g)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.getAll.projects"))
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
// @Param        id path int true "project ID"
// @Success      200  {object}  utils.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects/:id [get]
func (ph *ProjectHandler) GetByID(g *gin.Context) {

	dm, err := ph.Service.GetById(g, g.Param("id"))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.getByID.projects"))
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
// @Success      200  {object}  utils.Project
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects [post]
func (ph *ProjectHandler) Post(g *gin.Context) {

	dp := dto.Project{}
	err := g.BindJSON(&dp)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Post.bindJson.projects"))
		return
	}

	validate := validator.New()
	valErr := validate.StructExcept(dp, "ID")
	if valErr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valErr, "error.Post.validate.projects"))
		return
	}

	if len(dp.State) > 0 {
		err = dp.ValidateState()
		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Put.validateState.projects"))
			return
		}
	}

	dm, err := ph.Service.Insert(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.Post.insert.projects"))
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
// @Param        id path int true "project ID"
// @Param        state body utils.StateDTO true "Update a project's state"
// @Success      200  {object}  utils.Project
// @Failure      400  {object}	errors.ErrResponse
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

	err = g.BindJSON(&dp)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Patch.bindJson.projects"))
		return
	}

	validate := validator.New()
	valErr := validate.StructPartial(dp, "state")
	if valErr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valErr, "error.Patch.validate.projects"))
		return
	}

	err = dp.ValidateState()
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Put.validateState.projects"))
		return
	}

	dm, err := ph.Service.Update(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.Patch.update.projects"))
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
// @Param        id path int true "Project ID"
// @Param        project body dto.Project true "Update a project"
// @Success      200  {object}  utils.Project
// @Failure      400  {object}	errors.ErrResponse
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

	err = g.BindJSON(&dp)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Put.bindJson.projects"))
		return
	}

	if len(dp.State) > 0 {
		err = dp.ValidateState()
		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Put.validateState.projects"))
			return
		}
	}

	dm, err := ph.Service.Update(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(err, "error.Put.update.projects"))
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
// @Param        id path int true "Project ID"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}	errors.ErrResponse
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects/:id [delete]
func (ph *ProjectHandler) Delete(g *gin.Context) {
	dp := dto.Project{}

	i, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Delete.parseID.projects"))
		return
	}

	dp.ID = i

	_, err = ph.Service.Delete(g, dp.ToModel())

	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, errors.NewErrResponse(err, "error.Delete.delete.projects"))
		return
	}

	g.JSON(http.StatusOK, utils.Response{
		Message: "Project " + g.Param("id") + " deleted successfully",
	})

}
