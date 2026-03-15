package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/DenisHoliahaR/go-to-do/internal/task/service"
	"github.com/DenisHoliahaR/go-to-do/internal/transport"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service *service.Service
	logger  *slog.Logger
}

func NewTaskHandler(s *service.Service, l *slog.Logger) *handler {
	return &handler{
		service: s,
		logger: l,
	}
}

func (h *handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	proj, err := h.service.GetTaskById(r.Context(), id)

	preparedProj := TaskToTaskResponse(proj)
	transport.Write(w, http.StatusOK, preparedProj)
}

func (h *handler) GetTaskList(w http.ResponseWriter, r *http.Request) {
	projList, err := h.service.GetTaskList(r.Context())
	if err != nil {
		h.logger.Error("Failed to get tasks list", "error", err)
		http.Error(w, "Failed to get tasks list", http.StatusInternalServerError)
		return
	}

	preparedTaskList := TaskListToTaskListResponse(projList)
	transport.Write(w, http.StatusOK, preparedTaskList)
}

func (h *handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req CreateTaskRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid create Task request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := CreateTaskRequestToTask(req)
	createdTask, err := h.service.CreateTask(r.Context(), &data)
	if err != nil {
		h.logger.Error("Failed to update task", "error", err)
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
	res := TaskToTaskResponse(createdTask)

	transport.Write(w, http.StatusCreated, res)
}

func (h *handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req UpdateTaskRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid update Task request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := UpdateTaskRequestToTask(req)
	updatedTask, err := h.service.UpdateTask(r.Context(), &data, id); 
	if err != nil {
		h.logger.Error("Failed to update task", "error", err)
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
	res := TaskToTaskResponse(updatedTask)

	transport.Write(w, http.StatusOK, res)
}

func (h *handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(r.Context(), id); err != nil {
		h.logger.Error("Failed to delete task", "error", err)
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	transport.Write(w, http.StatusNoContent, nil)
}
