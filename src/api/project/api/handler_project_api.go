package api

import (
	//todo

	"net/http"
	"proyectos/src/api/errors"
	"proyectos/src/api/project/api/dto"
	"proyectos/src/api/project/domain"

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

	dm, err := ph.Service.Insert(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, errors.NewErrResponse(valerr))
		return
	}
	g.JSON(http.StatusOK, dto.FromModel(dm))

}
