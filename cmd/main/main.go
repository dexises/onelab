package main

import (
	"fmt"
	"log"
	"os"

	"onelab/internal/config"
	"onelab/internal/jsonlog"
	"onelab/internal/repository"
	"onelab/internal/service"
	"onelab/internal/transport"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Fatalf(fmt.Sprintf("Service shut down %v", run()))
}

func run() error {
	cfg := config.New()
	fmt.Println(cfg)

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	repo := repository.NewRepository()

	service := service.NewService(repo, logger)

	router := transport.NewRouter(service)

	err := transport.NewServer(cfg, logger, router)
	logger.PrintFatal(err, nil)
	return nil
}
