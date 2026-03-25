package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/DenisHoliahaR/go-to-do/internal/auth/service"
	"github.com/DenisHoliahaR/go-to-do/internal/transport"
)

type AuthService interface {
	SignUp(ctx context.Context, user *service.User) (int64, error)
	SignIn(ctx context.Context, user *service.User) (string, error)
}

type Handler struct {
	service AuthService
	logger  *slog.Logger
}

func NewAuthHandler(s *service.Service, l *slog.Logger) *Handler {
	return &Handler{
		service: s,
		logger:  l,
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req SignUpRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid sign up request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := SignUpRequestToUser(req)
	id, err := h.service.SignUp(r.Context(), &data)
	if err != nil {
		h.logger.Error("Failed to sign up new user", "error", err)
		http.Error(w, "Failed to sign up", http.StatusInternalServerError)
		return
	}

	res := UserToAuthResponse(id)

	transport.Write(w, http.StatusCreated, res)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req SignInRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid sign in request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := SignInRequestToUser(req)
	id, err := h.service.SignUp(r.Context(), &data)
	if err != nil {
		h.logger.Error("Failed to sign in new user", "error", err)
		http.Error(w, "Failed to sign in", http.StatusInternalServerError)
		return
	}

	res := UserToAuthResponse(id)

	transport.Write(w, http.StatusCreated, res)
}
