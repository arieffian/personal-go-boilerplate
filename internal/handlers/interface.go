package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthcheckService interface {
	HealthCheckHandler(c *fiber.Ctx) error
}

type CreateNewUserParams struct {
	Name string `json:"name" validate:"required"`
}

type UserService interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	CreateNewUser(c *fiber.Ctx) error
	// UpdateUser(c *fiber.Ctx) error
	// DeleteUser(c *fiber.Ctx) error
}
