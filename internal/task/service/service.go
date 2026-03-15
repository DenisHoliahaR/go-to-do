package service

import "context"

type TaskRepository interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	GetList(ctx context.Context) ([]*Task, error)
	GetByID(ctx context.Context, id int64) (*Task, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateTask(ctx context.Context, task *Task) (*Task, error) {
	return s.repo.Create(ctx, task)
}

func (s *Service) GetTaskById(ctx context.Context, id int64) (*Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetTaskList(ctx context.Context) ([]*Task, error) {
	return s.repo.GetList(ctx)
}

func (s *Service) UpdateTask(ctx context.Context, task *Task, id int64) (*Task, error) {
	return s.repo.Update(ctx, task)
}

func (s *Service) DeleteTask(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}