package routers

import (
	"github.com/arieffian/go-boilerplate/internal/config"
	"github.com/arieffian/go-boilerplate/internal/handlers"
	"github.com/arieffian/go-boilerplate/internal/middlewares"
	userRepository "github.com/arieffian/go-boilerplate/internal/repositories/users"
	database "github.com/arieffian/providers/pkg/db"
	"github.com/arieffian/providers/pkg/redis"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	healthcheck handlers.HealthcheckService
	users       handlers.UserService
	cfg         *config.Config
}

type NewRouterParams struct {
	Db    *database.DbInstance
	Redis redis.RedisService
	Cfg   *config.Config
}

func NewRouter(p NewRouterParams) (*Router, error) {
	userRepo := userRepository.NewUserRepository(userRepository.NewUserRepositoryParams{
		Db:    p.Db,
		Redis: p.Redis,
		Cfg:   p.Cfg,
	})

	healthcheckHandler := handlers.NewHealthCheckHandler()
	userHandler := handlers.NewUserHandler(handlers.NewUserHandlerParams{
		UserRepo: userRepo,
		Cfg:      p.Cfg,
	})

	return &Router{
		healthcheck: healthcheckHandler,
		users:       userHandler,
		cfg:         p.Cfg,
	}, nil
}

func (r *Router) RegisterRoutes(routes *fiber.App) {
	v1 := routes.Group("/api/v1")
	v1.Get("/healthcheck", r.healthcheck.HealthCheckHandler)

	users := v1.Group("/users").Use(middlewares.NewValidateAPIKey(r.cfg.ApiKey))
	users.Get("/", r.users.GetUsers)
	users.Get("/:id", r.users.GetUserById)
}
