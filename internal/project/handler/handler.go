package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/DenisHoliahaR/go-to-do/internal/project/service"
	"github.com/DenisHoliahaR/go-to-do/internal/transport"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service *service.Service
	logger  *slog.Logger
}

func NewProjectHandler(s *service.Service, l *slog.Logger) *handler {
	return &handler{
		service: s,
		logger: l,
	}
}

func (h *handler) GetProjectById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	proj, err := h.service.GetProjectById(r.Context(), id)

	preparedProj := ProjectToProjectResponse(proj)
	transport.Write(w, http.StatusOK, preparedProj)
}

func (h *handler) GetProjectList(w http.ResponseWriter, r *http.Request) {
	projList, err := h.service.GetProjectList(r.Context())
	if err != nil {
		h.logger.Error("Failed to get projects list", "error", err)
		http.Error(w, "Failed to get projects list", http.StatusInternalServerError)
		return
	}

	preparedProjectList := ProjectListToProjectListResponse(projList)
	transport.Write(w, http.StatusOK, preparedProjectList)
}

func (h *handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req CreateProjectRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid create Project request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := CreateProjectRequestToProject(req)
	createdProject, err := h.service.CreateProject(r.Context(), &data)
	if err != nil {
		h.logger.Error("Failed to update project", "error", err)
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}
	res := ProjectToProjectResponse(createdProject)

	transport.Write(w, http.StatusCreated, res)
}

func (h *handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var req UpdateProjectRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Error("Invalid update Project request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := UpdateProjectRequestToProject(req)
	updatedProject, err := h.service.UpdateProject(r.Context(), &data, id); 
	if err != nil {
		h.logger.Error("Failed to update project", "error", err)
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}
	res := ProjectToProjectResponse(updatedProject)

	transport.Write(w, http.StatusOK, res)
}

func (h *handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Request with invalid Id", "error", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProject(r.Context(), id); err != nil {
		h.logger.Error("Failed to delete project", "error", err)
		http.Error(w, "Failed to delete project", http.StatusInternalServerError)
		return
	}

	transport.Write(w, http.StatusNoContent, nil)
}
