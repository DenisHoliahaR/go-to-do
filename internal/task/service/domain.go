package service

import "time"

type Task struct {
	ID          int64
	Title       string
	Description string
	Status      TaskStatus
	ProjectID   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)
