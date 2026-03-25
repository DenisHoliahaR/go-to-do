package handler

import (
	"database/sql"
	"log/slog"

	"github.com/DenisHoliahaR/go-to-do/internal/task/repository"
	"github.com/DenisHoliahaR/go-to-do/internal/task/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RegisterHTTPEndpoints(r chi.Router, db *sql.DB, l *slog.Logger, t *jwtauth.JWTAuth) {
	repo := repository.NewTaskRepository(db)
	svc := service.NewTaskService(repo)
	h := NewTaskHandler(svc, l)

	r.Route("/tasks", func(r chi.Router) {
		r.Use(jwtauth.Verifier(t))
		r.Use(jwtauth.Authenticator(t))

		r.Post("/", h.CreateTask)
		r.Get("/", h.GetTaskList)
		r.Get("/{id}", h.GetTaskById)
		r.Put("/{id}", h.UpdateTask)
		r.Delete("/{id}", h.DeleteTask)
	})
}
