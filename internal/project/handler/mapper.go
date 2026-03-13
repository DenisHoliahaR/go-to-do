package handler

import domain "github.com/DenisHoliahaR/go-to-do/internal/project/service"

func ProjectToProjectResponse(project *domain.Project) ProjectResponse {
	return ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		OwnerID:     project.OwnerID,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func ProjectListToProjectListResponse(projects []*domain.Project) GetProjectListResponse {
	resp := GetProjectListResponse{
		Projects: make([]ProjectResponse, len(projects)),
	}

	for i, p := range projects {
		resp.Projects[i] = ProjectToProjectResponse(p)
	}

	return resp;
}

func CreateProjectRequestToProject(data CreateProjectRequest) domain.Project {
	return domain.Project{
		Name: data.Name,
		Description: data.Description,
		OwnerID: data.OwnerID,
	}
} 

func UpdateProjectRequestToProject(data UpdateProjectRequest) domain.Project {
	return domain.Project{
		Name: data.Name,
		Description: data.Description,
	}
} 