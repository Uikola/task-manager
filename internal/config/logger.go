package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/Uikola/task-manager/pkg/logger"
)

const (
	loggerConfigLevelKey = "LOG_LEVEL"
)

var (
	errLogLevelNotSet = errors.New("log level not set")
)

type Logger interface {
	Level() logger.LogLevel
}

type loggerConfig struct {
	level logger.LogLevel
}

func NewLoggerConfig() (Logger, error) {
	level := os.Getenv(loggerConfigLevelKey)
	levelInt, err := strconv.Atoi(level)
	if err != nil {
		return nil, errLogLevelNotSet
	}

	return &loggerConfig{
		level: logger.LogLevel(levelInt),
	}, nil
}

func (cfg *loggerConfig) Level() logger.LogLevel {
	return cfg.level
}
