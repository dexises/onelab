package main

import (
	"log"
	"os"

	"onelab/internal/config"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	config   config.Config
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	cfg := config.Load()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		config:   *cfg,
	}

	err := app.serve()
	errorLog.Fatal(err)
}
