package routers

import (
	"github.com/arieffian/go-boilerplate/internal/handlers"
	"github.com/arieffian/go-boilerplate/internal/pkg/redis"
	"github.com/arieffian/go-boilerplate/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Router struct {
	healthcheck handlers.HealthcheckService
	users       handlers.UserService
}

type NewRouterParams struct {
	Db    *gorm.DB
	Redis redis.RedisService
}

func NewRouter(p NewRouterParams) (*Router, error) {
	userRepo := repositories.NewUserRepository(repositories.NewUserRepositoryParams{
		Db:    p.Db,
		Redis: p.Redis,
	})

	healthcheckHandler := handlers.NewHealthCheckHandler()
	userHandler := handlers.NewUserHandler(handlers.NewUserHandlerParams{
		UserRepo: userRepo,
	})

	return &Router{
		healthcheck: healthcheckHandler,
		users:       userHandler,
	}, nil
}

func (r *Router) RegisterRoutes(routes *fiber.App) {
	v1 := routes.Group("/api/v1")
	v1.Get("/healthcheck", r.healthcheck.HealthCheckHandler)

	users := v1.Group("/users")
	users.Get("/", r.users.ListUsers)
	users.Get("/:id", r.users.GetUserById)
}
