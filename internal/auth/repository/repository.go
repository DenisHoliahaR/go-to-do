package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DenisHoliahaR/go-to-do/internal/auth/service"
)

type Repository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, user *service.User) (*service.User, error) {
	query := `
	INSERT INTO users(name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to insert user, err: %w", err)
	}

	return user, nil
}

func (r *Repository) GetById(ctx context.Context, id int64) (*service.User, error) {
	query := `
	SELECT * FROM users 
	WHERE id = $1
	`
	var user service.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to select user by id, err: %w", err)
	}

	return &user, nil
}
