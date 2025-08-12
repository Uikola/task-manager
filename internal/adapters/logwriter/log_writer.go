package logwriter

import "github.com/Uikola/task-manager/pkg/logger"

type LogWriter interface {
	// Start begins the log writer
	Start()

	// Stop gracefully shuts down the log writer
	Stop()

	// WriteLog sends a log entry for asynchronous processing
	WriteLog(level logger.LogLevel, message string, err error, fields logger.Fields)

	// IsRunning returns true if the log writer is currently active
	IsRunning() bool
}
