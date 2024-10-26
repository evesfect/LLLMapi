// internal/utils/logger/logger.go
package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
}

type defaultLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

func NewLogger() Logger {
	return &defaultLogger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
	}
}

func (l *defaultLogger) Info(args ...interface{}) {
	l.infoLogger.Println(args...)
}

func (l *defaultLogger) Error(args ...interface{}) {
	l.errorLogger.Println(args...)
}

func (l *defaultLogger) Fatal(args ...interface{}) {
	l.errorLogger.Fatal(args...)
}

func (l *defaultLogger) Debug(args ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		l.debugLogger.Println(args...)
	}
}
