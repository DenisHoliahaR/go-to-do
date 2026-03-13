package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DenisHoliahaR/go-to-do/internal/project/service"
)

type Repository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetList(ctx context.Context) ([]*service.Project, error) {
	query := `
	SELECT * FROM projects
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Projects repository list get %w", err)
	}
	defer rows.Close()

	var projects []*service.Project

	for rows.Next() {
		p := &service.Project{}
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.OwnerID,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan project %w", err)
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return projects, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*service.Project, error) {
	query := `
	SELECT * FROM projects
	WHERE id = $1
	`

	var proj service.Project
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&proj.ID,
		&proj.Name,
		&proj.Description,
		&proj.OwnerID,
		&proj.CreatedAt,
		&proj.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("project repository get %w", err)
	}

	return &proj, nil
}

func (r *Repository) Create(ctx context.Context, project *service.Project) (*service.Project, error) {
	query := `
	INSERT INTO projects(name, description, owner_id)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRowContext(ctx, query,
		project.Name,
		project.Description,
		project.OwnerID,
	).Scan(
		&project.ID,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("project repository create: %w", err)
	}

	return project, nil
}

func (r *Repository) Update(ctx context.Context, project *service.Project) (*service.Project, error) {
	query := `
	UPDATE projects
	SET name = $1, description = $2
	WHERE id = $3
	RETURNING id, owner_id, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		project.Name,
		project.Description,
	).Scan(
		&project.ID,
		&project.OwnerID,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("project repository update %w", err)
	}

	return project, nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM projects WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
