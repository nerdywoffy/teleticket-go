package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)
	return logger
}
