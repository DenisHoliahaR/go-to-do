package service

import "context"

type UserRepository interface {
	Create(ctx context.Context, project *User) (*User, error)
	GetList(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, project *User) (*User, error)
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateUser(ctx context.Context, project *User) (*User, error) {
	return s.repo.Create(ctx, project)
}

func (s *Service) GetUserById(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetUserList(ctx context.Context) ([]*User, error) {
	return s.repo.GetList(ctx)
}

func (s *Service) UpdateUser(ctx context.Context, project *User, id int64) (*User, error) {
	return s.repo.Update(ctx, project)
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
