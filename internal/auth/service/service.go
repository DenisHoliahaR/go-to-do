package service

import "github.com/DenisHoliahaR/go-to-do/internal/auth/repository"

type Service struct {
	repo repository.Repository
}

func NewAuthService(r repository.Repository) *Service {
	return &Service{
		repo: r,
	}
}