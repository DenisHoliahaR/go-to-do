package handler

import "github.com/DenisHoliahaR/go-to-do/internal/user/service"

func UserToUserResponse(user *service.User) UserResponse {
	return UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserListToUserListResponse(users []*service.User) GetUserListResponse {
	resp := &GetUserListResponse{
		Users: make([]UserResponse, len(users)),
	}
	for i, user := range users {
		resp.Users[i] = UserToUserResponse(user)
	}
	return *resp
}

func CreateUserRequestToUser(data CreateUserRequest) service.User {
	return service.User{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}
}

func UpdateUserRequestToUser(data UpdateUserRequest) service.User {
	return service.User{
		Name: data.Name,
	}
}
