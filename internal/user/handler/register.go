package handler

import (
	"database/sql"
	"log/slog"

	"github.com/DenisHoliahaR/go-to-do/internal/user/repository"
	"github.com/DenisHoliahaR/go-to-do/internal/user/service"
	"github.com/go-chi/chi/v5"
)

func RegisterHTTPEndpoints(r chi.Router, db *sql.DB, l *slog.Logger) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := NewUserHandler(svc, l)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/", h.GetUserList)
		r.Get("/{id}", h.GetUserById)
		r.Put("/{id}", h.UpdateUser)
		r.Delete("/{id}", h.DeleteUser)
	})
}
