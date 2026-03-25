package service

import (
	"context"
	"fmt"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

type ProjectRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetById(ctx context.Context, id int64) (*User, error)
}

type Service struct {
	repo ProjectRepository
	authToken *jwtauth.JWTAuth
}

func NewAuthService(r ProjectRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) SignUp(ctx context.Context, user *User) (int64, error) {
	ph, err := getHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = ph
	user, err = s.repo.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("Failed to create user, err: %s", err)
	}

	return user.ID, nil
}

func (s *Service) SignIn(ctx context.Context, user *User) (string, error) {
	repoUser, err := s.repo.GetById(ctx, user.ID)
	if err != nil {
		return "", fmt.Errorf("Failed to get user by ID, err: %w", err)
	}

	isUser := checkPassword(user.Password, repoUser.Password)
	if !isUser {
		return "", fmt.Errorf("Password mismatch")
	}

	claims := map[string]interface{}{"id": repoUser.ID, "email": repoUser.Email}
	_, tokenString, err := s.authToken.Encode(claims)
	if err != nil {
		return "", fmt.Errorf("Failed to encode claims, err: %w", err)
	}

	return tokenString, nil
}

func checkPassword(password, passwordHash string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return result == nil
}

func getHashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 25)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password, err: %w", err)
	}
	return string(passwordHash), nil
}
