package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthcheckService interface {
	HealthCheckHandler(c *fiber.Ctx) error
}

type UserService interface {
	ListUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	// CreateUser(c *fiber.Ctx) error
	// UpdateUser(c *fiber.Ctx) error
	// DeleteUser(c *fiber.Ctx) error
}
