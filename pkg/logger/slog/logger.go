package slog

import (
	"log/slog"
	"os"

	"github.com/Uikola/task-manager/pkg/logger"
)

type slogLogger struct {
	logger *slog.Logger
	level  logger.LogLevel
}

func NewLogger(level logger.LogLevel) logger.Logger {
	opts := &slog.HandlerOptions{
		Level: convertToSlogLevel(level),
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	log := slog.New(handler)

	return &slogLogger{
		logger: log,
		level:  level,
	}
}

func convertToSlogLevel(level logger.LogLevel) slog.Level {
	switch level {
	case logger.DebugLevel:
		return slog.LevelDebug
	case logger.InfoLevel:
		return slog.LevelInfo
	case logger.WarnLevel:
		return slog.LevelWarn
	case logger.ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func fieldsToArgs(fields ...logger.Fields) []any {
	var args []any
	for _, fieldMap := range fields {
		for key, value := range fieldMap {
			args = append(args, key, value)
		}
	}
	return args
}

func (l *slogLogger) Debug(msg string, fields ...logger.Fields) {
	if len(fields) > 0 {
		l.logger.Debug(msg, fieldsToArgs(fields...)...)
	} else {
		l.logger.Debug(msg)
	}
}

func (l *slogLogger) Info(msg string, fields ...logger.Fields) {
	if len(fields) > 0 {
		l.logger.Info(msg, fieldsToArgs(fields...)...)
	} else {
		l.logger.Info(msg)
	}
}

func (l *slogLogger) Warn(msg string, fields ...logger.Fields) {
	if len(fields) > 0 {
		l.logger.Warn(msg, fieldsToArgs(fields...)...)
	} else {
		l.logger.Warn(msg)
	}
}

func (l *slogLogger) Error(msg string, err error, fields ...logger.Fields) {
	args := fieldsToArgs(fields...)
	if err != nil {
		args = append(args, "error", err.Error())
	}

	if len(args) > 0 {
		l.logger.Error(msg, args...)
	} else {
		l.logger.Error(msg)
	}
}

func (l *slogLogger) WithFields(fields logger.Fields) logger.Logger {
	args := fieldsToArgs(fields)
	return &slogLogger{
		logger: l.logger.With(args...),
		level:  l.level,
	}
}

func (l *slogLogger) WithField(key string, value interface{}) logger.Logger {
	return &slogLogger{
		logger: l.logger.With(key, value),
		level:  l.level,
	}
}

// GetLevel returns the current log level
func (l *slogLogger) GetLevel() logger.LogLevel {
	return l.level
}
