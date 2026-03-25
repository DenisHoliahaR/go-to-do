package handler

import (
	"database/sql"
	"log/slog"

	"github.com/DenisHoliahaR/go-to-do/internal/auth/repository"
	"github.com/DenisHoliahaR/go-to-do/internal/auth/service"
	"github.com/go-chi/chi/v5"
)

func RegisterHTTPEndpoints(r chi.Router, db *sql.DB, l *slog.Logger) {
	repo := repository.NewAuthRepository(db)
	svc := service.NewAuthService(repo)
	h := NewAuthHandler(svc, l)

	r.Route("/", func(r chi.Router) {
		r.Post("/signup", h.SignUp)
		r.Post("/signin", h.SignIn)
	})

}