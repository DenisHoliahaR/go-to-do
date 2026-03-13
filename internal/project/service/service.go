package service

import "context"

type ProjectRepository interface {
	Create(ctx context.Context, project *Project) (*Project, error)
	GetList(ctx context.Context) ([]*Project, error)
	GetByID(ctx context.Context, id int64) (*Project, error)
	Update(ctx context.Context, project *Project) (*Project, error)
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	repo ProjectRepository
}

func NewProjectService(r ProjectRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	return s.repo.Create(ctx, project)
}

func (s *Service) GetProjectById(ctx context.Context, id int64) (*Project, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetProjectList(ctx context.Context) ([]*Project, error) {
	return s.repo.GetList(ctx)
}

func (s *Service) UpdateProject(ctx context.Context, project *Project, id int64) (*Project, error) {
	return s.repo.Update(ctx, project)
}

func (s *Service) DeleteProject(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}