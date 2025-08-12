package app

import (
	"fmt"

	"github.com/Uikola/task-manager/internal/config"
)

// serviceProvider acts as a centralized dependency container (DI).
// It lazily initializes and caches all shared services (singletons) and configuration required by the app.
type serviceProvider struct {
	httpConfig config.HTTP
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
