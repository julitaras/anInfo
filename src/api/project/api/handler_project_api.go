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

// Post ProjectCreator godoc
// @Summary      Add a project
// @Description  Add a project to the system
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        task body dto.Project true "Create a project"
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

	dm, err := ph.Service.Insert(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valerr))
		return
	}
	g.JSON(http.StatusOK, dto.FromModel(dm))

}

// Patch ProjectUpdater godoc
// @Summary      Update a project
// @Description  Update a project that is already in the system
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        task body dto.Project true "Update a project"
// @Success      200  {object}  dto.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects [patch]
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
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Invalid state",
		})
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
// @Param        task body dto.Project true "Create a project"
// @Success      200  {object}  dto.Project
// @Failure      422  {object}	errors.ErrResponse
// @Failure      500  {object}	errors.ErrResponse
// @Router       /projects [put]
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
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Unprocessable Entity",
		})
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
