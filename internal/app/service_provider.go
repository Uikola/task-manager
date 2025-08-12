package app

import (
	"fmt"

	"github.com/Uikola/task-manager/internal/usecase"
	"github.com/Uikola/task-manager/internal/usecase/task"

	"github.com/Uikola/task-manager/pkg/uuid"

	"github.com/Uikola/task-manager/internal/adapters/logwriter"
	"github.com/Uikola/task-manager/internal/adapters/logwriter/async"
	"github.com/Uikola/task-manager/pkg/closer"

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

	logger        logger.Logger
	uuidGenerator uuid.Generator

	taskRepository repository.TaskRepository

	taskUsecase usecase.TaskUsecase

	asyncLogWriter logwriter.LogWriter
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

func (s *serviceProvider) UUIDGenerator() uuid.Generator {
	if s.uuidGenerator == nil {
		s.uuidGenerator = uuid.NewGenerator()
	}

	return s.uuidGenerator
}

func (s *serviceProvider) TaskRepository() repository.TaskRepository {
	if s.taskRepository == nil {
		s.taskRepository = inmemory.NewTaskRepository()
	}

	return s.taskRepository
}

func (s *serviceProvider) TaskUsecase() usecase.TaskUsecase {
	if s.taskUsecase == nil {
		s.taskUsecase = task.NewUsecase(s.TaskRepository(), s.UUIDGenerator())
	}

	return s.taskUsecase
}

func (s *serviceProvider) AsyncLogWriter() logwriter.LogWriter {
	if s.asyncLogWriter == nil {
		s.asyncLogWriter = async.NewLogWriter(s.Logger(), 1000)

		closer.Add(func() error {
			s.asyncLogWriter.Stop()
			return nil
		})
	}

	return s.asyncLogWriter
}
