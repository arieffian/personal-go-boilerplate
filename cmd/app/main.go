package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/arieffian/go-boilerplate/internal/app"
	"github.com/arieffian/go-boilerplate/internal/config"
)

const (
	shutDownTimeout = 10 * time.Second
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config. %+v", err)
	}

	ctx := context.Background()

	server, err := app.NewServer(ctx, *cfg)
	if err != nil {
		log.Fatalf("failed to create the new server: %s\n", err)
	}

	go func() {
		if err := server.Listen(cfg.APIAddress); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeout)
	defer cancel()

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = server.Shutdown(ctx)

	fmt.Println("Fiber was successful shutdown.")
}
