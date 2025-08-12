package config

import (
	"fmt"
	"os"
)

func Load() error {
	if err := os.Setenv(httpConfigHostKey, "localhost"); err != nil {
		return fmt.Errorf("failed to set http host")
	}

	if err := os.Setenv(httpConfigPortKey, "8080"); err != nil {
		return fmt.Errorf("failed to set http port")
	}

	return nil
}
