package repositories

import "github.com/arieffian/go-boilerplate/internal/models"

type UserInterface interface {
	GetUserById(id string) (*models.User, error)
}
