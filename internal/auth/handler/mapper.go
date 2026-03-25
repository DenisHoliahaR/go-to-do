package handler

import "github.com/DenisHoliahaR/go-to-do/internal/auth/service"

func SignInRequestToUser(data SignInRequest) service.User {
	return service.User{
		Email: data.Email,
		Password: data.Password,
	}
}

func SignUpRequestToUser(data SignUpRequest) service.User {
	return service.User{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	}
}

func UserToAuthResponse(id int64) AuthResponse {
	return AuthResponse{
		ID: id,
	}
}