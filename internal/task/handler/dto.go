package handler

import (
	"time"

	"github.com/DenisHoliahaR/go-to-do/internal/task/service"
)

type CreateTaskRequest struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      service.TaskStatus `json:"status"`
	ProjectID   int64              `json:"project_id"`
}

type UpdateTaskRequest struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      service.TaskStatus `json:"status"`
}

type TaskResponse struct {
	ID          int64              `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      service.TaskStatus `json:"status"`
	ProjectID   int64              `json:"project_id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type GetTaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
