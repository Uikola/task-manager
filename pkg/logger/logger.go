package logger

// LogLevel represents the severity of a log message.
type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// String returns the string representation of the log level.
func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "unknown"
	}
}

// Fields represents structured logging fields.
type Fields map[string]interface{}

// Logger defines the logging interface
type Logger interface {
	// Basic logging methods
	Debug(msg string, fields ...Fields)
	Info(msg string, fields ...Fields)
	Warn(msg string, fields ...Fields)
	Error(msg string, err error, fields ...Fields)

	// Context-aware logging
	WithFields(fields Fields) Logger
	WithField(key string, value interface{}) Logger

	// Configuration
	GetLevel() LogLevel
}
