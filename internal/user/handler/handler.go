package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/DenisHoliahaR/go-to-do/internal/user/service"
	"github.com/DenisHoliahaR/go-to-do/internal/transport"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service *service.Service
	logger  *slog.Logger
}

func NewUserHandler(s *service.Service, l *slog.Logger) *handler {
	return &handler{
		service: s,
		logger: l,
	}
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserById(r.Context(), id)

	res := UserToUserResponse(user)
	transport.Write(w, http.StatusOK, res)
}

func (h *handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	userList, err := h.service.GetUserList(r.Context())
	if err != nil {
		h.logger.Error("Failed to get users list", "error", err)
		http.Error(w, "Failed to get users list", http.StatusInternalServerError)
		return
	}

	res := UserListToUserListResponse(userList)
	transport.Write(w, http.StatusOK, res)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid create Project request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := CreateUserRequestToUser(req)
	user, err := h.service.CreateUser(r.Context(), &data)
	if err != nil {
		h.logger.Error("Failed to update user", "error", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	res := UserToUserResponse(user)

	transport.Write(w, http.StatusCreated, res)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req UpdateUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid update Project request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := UpdateUserRequestToUser(req)
	user, err := h.service.UpdateUser(r.Context(), &data, id); 
	if err != nil {
		h.logger.Error("Failed to update user", "error", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	res := UserToUserResponse(user)

	transport.Write(w, http.StatusOK, res)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteUser(r.Context(), id); err != nil {
		h.logger.Error("Failed to delete user", "error", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	transport.Write(w, http.StatusNoContent, nil)
}