package async

import (
	"context"
	"sync"
	"time"

	"github.com/Uikola/task-manager/internal/adapters/logwriter"
	"github.com/Uikola/task-manager/pkg/logger"
)

type logEntry struct {
	Level   logger.LogLevel
	Message string
	Error   error
	Fields  logger.Fields
	Time    time.Time
}

// asyncLogWriter handles asynchronous logging through a channel.
type asyncLogWriter struct {
	logger     logger.Logger
	logChan    chan logEntry
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	bufferSize int
	running    bool
	mu         sync.RWMutex
}

// NewLogWriter creates a new asynchronous log writer.
func NewLogWriter(logger logger.Logger, bufferSize int) logwriter.LogWriter {
	ctx, cancel := context.WithCancel(context.Background())

	return &asyncLogWriter{
		logger:     logger,
		logChan:    make(chan logEntry, bufferSize),
		ctx:        ctx,
		cancel:     cancel,
		bufferSize: bufferSize,
		running:    false,
	}
}

// Start begins the asynchronous logging process.
func (w *asyncLogWriter) Start() {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.running {
		return
	}

	w.logger.Info("starting async log writer")

	w.running = true
	w.wg.Add(1)
	go w.processLogs()
}

// processLogs is the main loop that processes log entries from the channel
func (w *asyncLogWriter) processLogs() {
	defer w.wg.Done()

	for {
		select {
		case entry, ok := <-w.logChan:
			if !ok {
				return
			}
			w.writeLogEntry(entry)

		case <-w.ctx.Done():
			w.emptyChannel()
			return
		}
	}
}

// emptyChannel processes any remaining log entries in the channel
func (w *asyncLogWriter) emptyChannel() {
	for {
		select {
		case entry, ok := <-w.logChan:
			if !ok {
				return
			}
			w.writeLogEntry(entry)
		default:
			return
		}
	}
}

// writeLogEntry writes a single log entry using the underlying logger
func (w *asyncLogWriter) writeLogEntry(entry logEntry) {
	switch entry.Level {
	case logger.DebugLevel:
		w.logger.Debug(entry.Message, entry.Fields)
	case logger.InfoLevel:
		w.logger.Info(entry.Message, entry.Fields)
	case logger.WarnLevel:
		w.logger.Warn(entry.Message, entry.Fields)
	case logger.ErrorLevel:
		w.logger.Error(entry.Message, entry.Error, entry.Fields)
	}
}

// Stop gracefully shuts down the async log writer.
func (w *asyncLogWriter) Stop() {
	w.mu.Lock()
	if !w.running {
		w.mu.Unlock()
		return
	}
	w.running = false
	w.mu.Unlock()

	w.logger.Info("stopping async log writer")

	w.cancel()
	close(w.logChan)

	w.wg.Wait()
}

// WriteLog sends a log entry for asynchronous processing.
func (w *asyncLogWriter) WriteLog(level logger.LogLevel, message string, err error, fields logger.Fields) {
	if !w.IsRunning() {
		return
	}

	entry := logEntry{
		Level:   level,
		Message: message,
		Error:   err,
		Fields:  fields,
		Time:    time.Now(),
	}

	select {
	case w.logChan <- entry:
	case <-w.ctx.Done():
	}
}

// IsRunning returns true if the log writer is currently active
func (w *asyncLogWriter) IsRunning() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.running
}
