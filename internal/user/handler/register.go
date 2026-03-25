package handler

import (
	"database/sql"
	"log/slog"

	"github.com/DenisHoliahaR/go-to-do/internal/user/repository"
	"github.com/DenisHoliahaR/go-to-do/internal/user/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RegisterHTTPEndpoints(r chi.Router, db *sql.DB, l *slog.Logger, t *jwtauth.JWTAuth) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := NewUserHandler(svc, l)

	r.Route("/users", func(r chi.Router) {
		r.Use(jwtauth.Verifier(t))
		r.Use(jwtauth.Authenticator(t))
		
		r.Get("/", h.GetUserList)
		r.Get("/{id}", h.GetUserById)
		r.Put("/{id}", h.UpdateUser)
		r.Delete("/{id}", h.DeleteUser)
	})
}
