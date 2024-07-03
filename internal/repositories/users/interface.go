package users

import (
	"context"

	"github.com/arieffian/go-boilerplate/internal/models"
)

type GetUserByIdParams struct {
	UserId string
}

type GetUserByIdResponse struct {
	User models.User
}

type GetUsersParams struct {
	Limit  int
	Offset int
}

type GetUsersResponse struct {
	Users []models.User
}

type CreateNewUserParams struct {
	Name string
}

type CreateNewUserResponse struct {
	User models.User
}

type UserRepository interface {
	GetUserById(ctx context.Context, p GetUserByIdParams) (*GetUserByIdResponse, error)
	GetUsers(ctx context.Context, p GetUsersParams) (*GetUsersResponse, error)
	CreateNewUser(ctx context.Context, p CreateNewUserParams) (*CreateNewUserResponse, error)
}
