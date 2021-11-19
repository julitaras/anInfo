package api

import (
	//todo

	"fmt"
	"net/http"
	"proyectos/src/api/thing/api/dto"
	"proyectos/src/api/thing/domain"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

//ThingHandler handler
type ThingHandler struct {
	domain.Service
}

//Get handler
func (dh *ThingHandler) Get(g *gin.Context) {

	dt := dto.Thing{}
	g.BindQuery(&dt)

	validate := validator.New()
	valerr := validate.Var(dt.ID, "gt=0")
	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusConflict, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Debe indicar un ID",
		})
		return
	}

	dd, err := dh.Service.Get(g, dt.ID)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusConflict, ErrResponse{
			Error:   err.Error(),
			Message: "Error obteniendo datos.",
		})
	} else {
		g.JSON(http.StatusOK, dd)
	}
}

//Post handler
func (dh *ThingHandler) Post(g *gin.Context) {

	dt := dto.Thing{}
	g.BindJSON(&dt)

	validate := validator.New()

	valerr := validate.StructExcept(dt, "ID")

	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Unprocessable Entity",
		})
		return
	}

	dm, err := dh.Service.Insert(g, dt.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot Save",
		})
	} else {
		g.JSON(http.StatusOK, dm)
	}
}

//Put handler
func (dh *ThingHandler) Put(g *gin.Context) {

	dt := dto.Thing{}
	g.BindJSON(&dt)

	validate := validator.New()

	valerr := validate.Struct(dt)

	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Unprocessable Entity",
		})
		return
	}

	dm, err := dh.Service.Insert(g, dt.ToModel())
	if err != nil {
		g.AbortWithStatusJSON(http.StatusConflict, ErrResponse{
			Error:   err.Error(),
			Message: "Cannot Save",
		})
	} else {
		g.JSON(http.StatusOK, dm)
	}
}

//Delete handler
func (dh *ThingHandler) Delete(g *gin.Context) {

	dt := dto.Thing{}
	g.BindQuery(&dt)

	validate := validator.New()
	valerr := validate.Var(dt.ID, "gt=0")
	if valerr != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, ErrResponse{
			Error:   getValErr(valerr.(validator.ValidationErrors)),
			Message: "Debe indicar un ID",
		})
		return
	}

	dd, err := dh.Service.Delete(g, dt.ID)

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
