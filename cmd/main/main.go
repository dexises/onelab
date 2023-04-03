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
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	config   config.Config
}

func main() {
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	// cfg := config.Load()

	log.Fatalln(fmt.Sprintf("Service shut down %v", run()))
}

func run() error {
	cfg := config.New()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	repo := repository.NewRepository()

	service := service.NewService(repo)

	router := transport.NewRouter(service)

	return nil
}
