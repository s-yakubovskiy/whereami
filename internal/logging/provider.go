package logging

import (
	"github.com/s-yakubovskiy/whereami/internal/config"
	"github.com/s-yakubovskiy/whereami/pkg/shudralogs"
)

// ProvideLogger configures and returns a zerolog logger based on the application's configuration.
func ProvideLogger(cfg *config.AppConfig) shudralogs.Logger {
	return shudralogs.NewLoggerWithConfig(&shudralogs.Config{
		Format: cfg.Logging.Format,
		Output: cfg.Logging.Output,
		Level:  shudralogs.LogLevelFromString(cfg.Logging.Level),
	})
}
