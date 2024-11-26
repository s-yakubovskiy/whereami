package logging

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/s-yakubovskiy/whereami/internal/config"
)

func logLevelFromString(logLevel string) zerolog.Level {
	level := strings.ToLower(logLevel)
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}

// ProvideLogger configures and returns a zerolog logger based on the application's configuration.
func ProvideLogger(cfg *config.AppConfig) Logger {
	var logger zerolog.Logger

	// Set the log level
	level := logLevelFromString(strings.ToLower(cfg.Logging.Level))
	zerolog.SetGlobalLevel(level)

	// Set the log format
	if strings.ToLower(cfg.Logging.Format) == "json" {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger() // JSON is default, but setting explicitly
	} else {
		// Configure the ConsoleWriter explicitly
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}

	// Set the log output
	switch cfg.Logging.Output {
	case "stdout":
		// Do nothing since it's already the default
	case "stderr":
		logger = logger.Output(os.Stderr)
	default:
		file, err := os.OpenFile(cfg.Logging.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logger = logger.Output(os.Stdout) // Fallback to stdout if file cannot be opened
			logger.Error().Err(err).Msg("Failed to open log file, falling back to stdout")
		} else {
			logger = logger.Output(file)
		}
	}

	return &ZerologLogger{logger: logger}
}
