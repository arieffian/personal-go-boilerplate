package handlers

import (
	"errors"

	"github.com/arieffian/go-boilerplate/internal/config"
	"github.com/arieffian/go-boilerplate/internal/pkg/generated"
	userRepository "github.com/arieffian/go-boilerplate/internal/repositories/users"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userHandler struct {
	userRepo userRepository.UserInterface
	cfg      *config.Config
}

var _ UserService = (*userHandler)(nil)

type NewUserHandlerParams struct {
	UserRepo userRepository.UserInterface
	Cfg      *config.Config
}

func (h userHandler) ListUsers(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func (h userHandler) GetUserById(c *fiber.Ctx) error {
	user, err := h.userRepo.GetUserById(c.Context(), userRepository.GetUserByIdParams{
		UserId: c.Params("id"),
	})

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.GetUserByIdResponse{
			Code:    int32(status),
			Message: "OK",
			Data: &generated.User{
				Id:   user.User.ID,
				Name: user.User.Name,
			},
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.GetUserByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.User{
			Id:   user.User.ID,
			Name: user.User.Name,
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func NewUserHandler(p NewUserHandlerParams) *userHandler {
	return &userHandler{
		userRepo: p.UserRepo,
		cfg:      p.Cfg,
	}
}
