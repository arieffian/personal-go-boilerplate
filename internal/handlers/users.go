package handlers

import (
	"errors"
	"strconv"

	"github.com/arieffian/go-boilerplate/internal/config"
	"github.com/arieffian/go-boilerplate/internal/pkg/generated"
	userRepository "github.com/arieffian/go-boilerplate/internal/repositories/users"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userHandler struct {
	userRepo userRepository.UserRepository
	cfg      *config.Config
}

var _ UserService = (*userHandler)(nil)

type NewUserHandlerParams struct {
	UserRepo userRepository.UserRepository
	Cfg      *config.Config
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	pPage := c.Params("page")
	page, err := strconv.Atoi(pPage)

	if err != nil {
		status := fiber.StatusBadRequest
		response := generated.GetUsersResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	if page <= 0 {
		page = 1
	}

	usersResult, err := h.userRepo.GetUsers(c.Context(), userRepository.GetUsersParams{
		Limit:  h.cfg.DefaultPageSize,
		Offset: (page - 1) * h.cfg.DefaultPageSize,
	})

	if err != nil {
		status := fiber.StatusInternalServerError
		response := generated.GetUsersResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	var users []generated.User
	for _, user := range usersResult.Users {
		users = append(users, generated.User{
			Id:   user.ID,
			Name: user.Name,
		})
	}

	response := generated.GetUsersResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    &users,
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *userHandler) GetUserById(c *fiber.Ctx) error {
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

func NewUserHandler(p NewUserHandlerParams) UserService {
	return &userHandler{
		userRepo: p.UserRepo,
		cfg:      p.Cfg,
	}
}
