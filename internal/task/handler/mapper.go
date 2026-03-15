package handler

import "github.com/DenisHoliahaR/go-to-do/internal/task/service"

func TaskToTaskResponse(task *service.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		ProjectID:   task.ProjectID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func TaskListToTaskListResponse(tasks []*service.Task) GetTaskListResponse {
	resp := GetTaskListResponse{
		Tasks: make([]TaskResponse, len(tasks)),
	}

	for i, p := range tasks {
		resp.Tasks[i] = TaskToTaskResponse(p)
	}

	return resp
}

func CreateTaskRequestToTask(data CreateTaskRequest) service.Task {
	return service.Task{
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		ProjectID:   data.ProjectID,
	}
}

func UpdateTaskRequestToTask(data UpdateTaskRequest) service.Task {
	return service.Task{
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
	}
}
