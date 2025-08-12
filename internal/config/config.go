package config

import (
	"fmt"
	"os"
)

func Load() error {
	if err := os.Setenv(httpConfigHostKey, "0.0.0.0"); err != nil {
		return fmt.Errorf("failed to set http host")
	}

	if err := os.Setenv(httpConfigPortKey, "8080"); err != nil {
		return fmt.Errorf("failed to set http port")
	}

	if err := os.Setenv(loggerConfigLevelKey, "0"); err != nil {
		return fmt.Errorf("failed to set log level")
	}

	return nil
}
