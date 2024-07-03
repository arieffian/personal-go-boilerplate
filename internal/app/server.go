package app

import (
	"context"
	"log"

	"github.com/arieffian/go-boilerplate/internal/config"
	"github.com/arieffian/go-boilerplate/internal/routers"
	database "github.com/arieffian/providers/pkg/db"
	"github.com/arieffian/providers/pkg/redis"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Fiber *fiber.App
}

func NewServer(ctx context.Context, cfg *config.Config) (*Server, error) {

	db := database.NewDbManager(database.DbConfig{
		WriteDsn: cfg.DbMasterConnectionString,
		ReadDsn:  cfg.DbReplicaConnectionString,
	})

	dbClient, err := db.CreateDbClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database. %+v", err)
	}

	redis := redis.NewRedisConnection(redis.RedisConfig{
		Host: cfg.RedisHost,
		Port: cfg.RedisPort,
	})

	app := fiber.New()

	api, err := routers.NewRouter(routers.NewRouterParams{
		Db:    dbClient,
		Redis: redis,
		Cfg:   cfg,
	})

	if err != nil {
		return nil, err
	}

	api.RegisterRoutes(app)

	return &Server{
		Fiber: app,
	}, nil

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Fiber.ShutdownWithContext(ctx)
}

func (s *Server) Listen(addr string) error {
	return s.Fiber.Listen(addr)
}
