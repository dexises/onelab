package logger

import (
	"log"
	"os"
)

type ILogger interface {
	Printf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type Logger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func New() *Logger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	return &Logger{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

func (l *Logger) InfoLog(format string, v ...interface{}) {
	l.infoLog.Printf(format, v...)
}

func (l *Logger) ErrorLog(format string, v ...interface{}) {
	l.errorLog.Printf(format, v...)
}
