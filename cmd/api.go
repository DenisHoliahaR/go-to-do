package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/DenisHoliahaR/go-to-do/internal/infrastructure/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	projectH "github.com/DenisHoliahaR/go-to-do/internal/project/handler"
	projectR "github.com/DenisHoliahaR/go-to-do/internal/project/repository"
	projectS "github.com/DenisHoliahaR/go-to-do/internal/project/service"

	taskH "github.com/DenisHoliahaR/go-to-do/internal/task/handler"
	taskR "github.com/DenisHoliahaR/go-to-do/internal/task/repository"
	taskS "github.com/DenisHoliahaR/go-to-do/internal/task/service"

	userH "github.com/DenisHoliahaR/go-to-do/internal/user/handler"
	userR "github.com/DenisHoliahaR/go-to-do/internal/user/repository"
	userS "github.com/DenisHoliahaR/go-to-do/internal/user/service"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	db := postgres.NewPostgresDB(app.config.db.dsn)

	if db == nil {
		app.logger.Error("Failed to connect to database")
		return nil
	}

	projectRepository := projectR.NewProjectRepository(db)
	projectService := projectS.NewProjectService(projectRepository)
	projectHandler := projectH.NewProjectHandler(projectService, app.logger)
	r.Route("/projects", func(r chi.Router) {
		r.Post("/", projectHandler.CreateProject)
		r.Get("/", projectHandler.GetProjectList)
		r.Get("/{id}", projectHandler.GetProjectById)
		r.Put("/{id}", projectHandler.UpdateProject)
		r.Delete("/{id}", projectHandler.DeleteProject)
	})

	taskRepository := taskR.NewTaskRepository(db)
	taskService := taskS.NewTaskService(taskRepository)
	taskHandler := taskH.NewTaskHandler(taskService, app.logger)
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskHandler.CreateTask)
		r.Get("/", taskHandler.GetTaskList)
		r.Get("/{id}", taskHandler.GetTaskById)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
	})

	userRepository := userR.NewUserRepository(db)
	userService := userS.NewUserService(userRepository)
	userHandler := userH.NewUserHandler(userService, app.logger)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetUserList)
		r.Get("/{id}", userHandler.GetUserById)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Info("Server has started", "address", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	logger *slog.Logger
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
