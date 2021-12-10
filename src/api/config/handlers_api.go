package config

import (
	"gorm.io/gorm"
	projectApi "proyectos/src/api/project/api"
	projectDomain "proyectos/src/api/project/domain"
	projectRepository "proyectos/src/api/project/repository"
	projectService "proyectos/src/api/project/service"
	taskApi "proyectos/src/api/task/api"
	taskDomain "proyectos/src/api/task/domain"
	taskRepository "proyectos/src/api/task/repository"
	taskService "proyectos/src/api/task/service"
)

//AddHandlers routes
func (r *SRV) AddHandlers(db *gorm.DB) *SRV {
	tr := taskRepository.NewTaskRepository(db)
	ts := taskService.NewTaskService(tr)
	pr := projectRepository.NewProjectRepository(db)
	ps := projectService.NewProjectService(pr)

	r = AddTaskHandler(r, ts)
	r = AddProjectHandler(r, ps)
	return r
}

//AddTaskHandler routes set
func AddTaskHandler(r *SRV, ds taskDomain.Service) *SRV {
	taskHandler := &taskApi.TaskHandler{
		Service: ds,
	}

	r.POST("/tasks", taskHandler.Post)
	r.DELETE("/tasks/:id", taskHandler.Delete)

	return r
}

//AddProjectHandler routes set
func AddProjectHandler(r *SRV, ds projectDomain.Service) *SRV {
	projectHandler := &projectApi.ProjectHandler{
		Service: ds,
	}

	r.POST("/projects", projectHandler.Post)
	r.PATCH("/projects/:id/state", projectHandler.Patch)
	r.PUT("/projects/:id", projectHandler.Put)

	return r
}
