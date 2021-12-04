package api

import (
	//todo

	"fmt"
	"net/http"
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
	g.BindJSON(&dp)

	validate := validator.New()

	valerr := validate.StructExcept(dp, "ID")

	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Unprocessable Entity",
		})
		return
	}

	dm, err := ph.Service.Insert(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot Save",
		})
	} else {
		g.JSON(http.StatusOK, dm)
	}
}

func (ph *ProjectHandler) Patch(g *gin.Context) {

	dp := dto.Project{}
	g.BindJSON(&dp)

	validate := validator.New()

	valerr := validate.StructPartial(dp, "ID", "State")

	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
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
	} else {
		g.JSON(http.StatusOK, dm)
	}
}

func getValErr(e validator.ValidationErrors) string {
	var ee string
	for _, err := range e {
		ee = fmt.Sprintf("%s Field: %s, Type: %s, Error Value: %s /n", ee, err.Field(), err.Type(), err.Value())
	}
	return ee
}
