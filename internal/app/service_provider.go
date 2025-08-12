package app

import (
	"fmt"

	"github.com/Uikola/task-manager/internal/adapters/repository"
	"github.com/Uikola/task-manager/internal/adapters/repository/inmemory"

	"github.com/Uikola/task-manager/pkg/logger"
	"github.com/Uikola/task-manager/pkg/logger/slog"

	"github.com/Uikola/task-manager/internal/config"
)

// serviceProvider acts as a centralized dependency container (DI).
// It lazily initializes and caches all shared services (singletons) and configuration required by the app.
type serviceProvider struct {
	httpConfig   config.HTTP
	loggerConfig config.Logger

	logger logger.Logger

	taskRepository repository.TaskRepository
}

// newServiceProvider returns a new, empty DI provider.
// All dependencies are constructed lazily in their respective getters.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTP {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get http config: %w", err))
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) LoggerConfig() config.Logger {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get logger config: %w", err))
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serviceProvider) Logger() logger.Logger {
	if s.logger == nil {
		s.logger = slog.NewLogger(s.LoggerConfig().Level())
	}

	return s.logger
}

func (s *serviceProvider) TaskRepository() repository.TaskRepository {
	if s.taskRepository == nil {
		s.taskRepository = inmemory.NewTaskRepository()
	}

	return s.taskRepository
}
