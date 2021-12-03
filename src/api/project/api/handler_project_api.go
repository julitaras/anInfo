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

// Get handler
func (ph *ProjectHandler) Get(g *gin.Context) {

	dp := dto.Project{}
	g.BindQuery(&dp)

	validate := validator.New()
	valerr := validate.Var(dp.Code, "gt=0")
	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusConflict, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Debe indicar un código para el proyecto, que sea un número mayor que cero",
		})
		return
	}

	dd, err := ph.Service.Get(g, dp.Code)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusConflict, ErrResponse{
			Error:   err.Error(),
			Message: "Error obteniendo datos.",
		})
	} else {
		g.JSON(http.StatusOK, dd)
	}
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

// Put handler
func (ph *ProjectHandler) Put(g *gin.Context) {

	dp := dto.Project{}
	g.BindJSON(&dp)

	validate := validator.New()

	valerr := validate.Struct(dp)

	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Unprocessable Entity",
		})
		return
	}

	dm, err := ph.Service.Insert(g, dp.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusConflict, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot Save",
		})
	} else {
		g.JSON(http.StatusOK, dm)
	}
}

// Delete handler
func (ph *ProjectHandler) Delete(g *gin.Context) {

	dp := dto.Project{}
	g.BindQuery(&dp)

	validate := validator.New()
	valerr := validate.Var(dp.Code, "gt=0")
	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Debe indicar un ID",
		})
		return
	}

	dd, err := ph.Service.Delete(g, dp.Code)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusConflict, ErrResponse{
			Error:   err.Error(),
			Message: "Error obteniendo datos.",
		})
		return
	}
	g.JSON(http.StatusOK, dd)

}

func getValErr(e validator.ValidationErrors) string {
	var ee string
	for _, err := range e {
		ee = fmt.Sprintf("%s Field: %s, Type: %s, Error Value: %s /n", ee, err.Field(), err.Type(), err.Value())
	}
	return ee
}
