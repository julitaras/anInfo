package config

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	_ "proyectos/src/api/docs"
	projectApi "proyectos/src/api/project/api"
	projectDomain "proyectos/src/api/project/domain"
	projectRepository "proyectos/src/api/project/repository"
	projectService "proyectos/src/api/project/service"
	taskApi "proyectos/src/api/task/api"
	taskDomain "proyectos/src/api/task/domain"
	taskRepository "proyectos/src/api/task/repository"
	taskService "proyectos/src/api/task/service"
)

// @title           PSA Projects API
// @version         1.0
// @description     This API gives access to the projects module.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://squad14-2c-2021.herokuapp.com/


//AddHandlers routes
func (r *SRV) AddHandlers(db *gorm.DB) *SRV {
	tr := taskRepository.NewTaskRepository(db)
	ts := taskService.NewTaskService(tr)
	pr := projectRepository.NewProjectRepository(db)
	ps := projectService.NewProjectService(pr)

	r = AddTaskHandler(r, ts)
	r = AddProjectHandler(r, ps)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

//AddTaskHandler routes set
func AddTaskHandler(r *SRV, ds taskDomain.Service) *SRV {
	taskHandler := &taskApi.TaskHandler{
		Service: ds,
	}

	r.POST("/tasks", taskHandler.Post)
	r.PUT("/tasks/:id", taskHandler.Put)
	r.DELETE("/tasks/:id", taskHandler.Delete)
	r.GET("/tasks", taskHandler.GetAll)
	r.GET("/tasks/:id", taskHandler.GetByID)

	return r
}

//AddProjectHandler routes set
func AddProjectHandler(r *SRV, ds projectDomain.Service) *SRV {
	projectHandler := &projectApi.ProjectHandler{
		Service: ds,
	}

	r.GET("/projects", projectHandler.GetAll)
	r.GET("/projects/:id", projectHandler.GetByID)
	r.POST("/projects", projectHandler.Post)
	r.PATCH("/projects/:id/state", projectHandler.Patch)
	r.PUT("/projects/:id", projectHandler.Put)
	r.DELETE("/projects/:id", projectHandler.Delete)

	return r
}
