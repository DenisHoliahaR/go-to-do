package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DenisHoliahaR/go-to-do/internal/task/service"
)

type Repository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetList(ctx context.Context) ([]*service.Task, error) {
	query := `
	SELECT * FROM tasks
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Tasks repository list get %w", err)
	}
	defer rows.Close()

	var tasks []*service.Task

	for rows.Next() {
		task := &service.Task{}
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.ProjectID,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan task %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tasks, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*service.Task, error) {
	query := `
	SELECT * FROM tasks
	WHERE id = $1
	`

	var task service.Task
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.ProjectID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("task repository get %w", err)
	}

	return &task, nil
}

func (r *Repository) Create(ctx context.Context, task *service.Task) (*service.Task, error) {
	query := `
	INSERT INTO tasks(title, description, status, project_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRowContext(ctx, query,
		task.Title,
		task.Description,
		task.Status,
		task.ProjectID,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("task repository create: %w", err)
	}

	return task, nil
}

func (r *Repository) Update(ctx context.Context, task *service.Task) (*service.Task, error) {
	query := `
	UPDATE tasks
	SET title = $1, description = $2, status = $3
	WHERE id = $4
	RETURNING id, project_id, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		task.Title,
		task.Description,
		task.Status,
		task.ID,
	).Scan(
		&task.ID,
		&task.ProjectID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("task repository update %w", err)
	}

	return task, nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
