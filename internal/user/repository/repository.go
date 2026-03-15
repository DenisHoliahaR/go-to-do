package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DenisHoliahaR/go-to-do/internal/user/service"
)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetList(ctx context.Context) ([]*service.User, error) {
	query := `
	SELECT * FROM users
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Users repository list get %w", err)
	}
	defer rows.Close()

	var users []*service.User

	for rows.Next() {
		p := &service.User{}
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Email,
			&p.Password,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan user %w", err)
		}
		users = append(users, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*service.User, error) {
	query := `
	SELECT * FROM users
	WHERE id = $1
	`

	var user service.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("user repository get %w", err)
	}

	return &user, nil
}

func (r *Repository) Create(ctx context.Context, user *service.User) (*service.User, error) {
	query := `
	INSERT INTO users(name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRowContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("user repository create: %w", err)
	}

	return user, nil
}

func (r *Repository) Update(ctx context.Context, user *service.User) (*service.User, error) {
	query := `
	UPDATE users
	SET name = $1,
	WHERE id = $2
	RETURNING id, email, password, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		user.Name,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("user repository update %w", err)
	}

	return user, nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
