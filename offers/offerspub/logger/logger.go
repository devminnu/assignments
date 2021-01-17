package logger

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

//Init init logger
func Init() {
	logger = log.New()
	logger.Formatter = &log.JSONFormatter{}
	logPath := os.Getenv("APP_LOG_PATH")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.SetLevel(log.InfoLevel)
}

func Get() *log.Logger {
	if logger == nil {
		logrus.Error("LOGGER_NOT_INITIALIZED")
		panic(errors.New("LOGGER_NOT_INITIALIZED"))
	}
	return logger
}
