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

// Post handler
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
