package handler

import (
	"database/sql"
	"log/slog"

	"github.com/DenisHoliahaR/go-to-do/internal/project/repository"
	"github.com/DenisHoliahaR/go-to-do/internal/project/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RegisterHTTPEndpoints(r chi.Router, db *sql.DB, l *slog.Logger, t *jwtauth.JWTAuth) {
	repo := repository.NewProjectRepository(db)
	svc := service.NewProjectService(repo)
	h := NewProjectHandler(svc, l)

	r.Route("/projects", func(r chi.Router) {
		r.Use(jwtauth.Verifier(t))
		r.Use(jwtauth.Authenticator(t))

		r.Post("/", h.CreateProject)
		r.Get("/", h.GetProjectList)
		r.Get("/{id}", h.GetProjectByID)
		r.Put("/{id}", h.UpdateProject)
		r.Delete("/{id}", h.DeleteProject)
	})
}