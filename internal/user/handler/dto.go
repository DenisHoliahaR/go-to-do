package handler

import "time"

type UserResponse struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password"`
}

type GetUserListResponse struct {
	Users []UserResponse `json:"users"` 
}
