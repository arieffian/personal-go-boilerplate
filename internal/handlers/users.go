package handlers

import (
	"errors"

	"github.com/arieffian/go-boilerplate/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userHandler struct {
	userRepo repositories.UserInterface
}

var _ UserService = (*userHandler)(nil)

type NewUserHandlerParams struct {
	UserRepo repositories.UserInterface
}

func (h userHandler) ListUsers(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func (h userHandler) GetUserById(c *fiber.Ctx) error {
	user, err := h.userRepo.GetUserById(c.Params("id"))

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := Response{
			Status:  status,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(response.Status).JSON(response)
	}

	response := Response{
		Status:  fiber.StatusOK,
		Message: "OK",
		Data:    user,
	}
	return c.Status(response.Status).JSON(response)
}

func NewUserHandler(p NewUserHandlerParams) *userHandler {
	return &userHandler{
		userRepo: p.UserRepo,
	}
}
