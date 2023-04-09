package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"onelab/internal/config"
	"onelab/internal/jsonlog"
	"onelab/internal/repository"
	"onelab/internal/service"
	"onelab/internal/transport"
	"onelab/internal/transport/http/handler"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Fatalf(fmt.Sprintf("Service shut down %v", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	repo := repository.NewRepository(ctx, cfg)

	service := service.NewService(repo, logger)

	h := handler.NewManager(cfg, service)

	err := transport.NewServer(cfg, h)
	logger.PrintFatal(err, nil)
	return nil
}
