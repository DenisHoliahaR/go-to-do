package handler

import (
	"log/slog"

	"github.com/DenisHoliahaR/go-to-do/internal/auth/service"
)

type Handler struct {
	service *service.Service
	logger *slog.Logger
}

func NewAuthHandler(s *service.Service, l *slog.Logger) *Handler {
	return &Handler{
		service: s,
		logger: l,
	}
}
