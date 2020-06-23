package project

type ProjectService interface {
	GetAll() []*Project
	GetProjectByName(projectName string) *Project
}

var projectService = &projectService{}

func ProjectServiceInstance(repo ProjectRepository) ProjectService {
	projectServiceInstance.repo = repo
	return projectServiceInstance
}

type projectService struct {
	repo ProjectRepository
}
