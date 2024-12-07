package shudralogs

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Config holds the logger configuration.
type Config struct {
	Level  zerolog.Level
	Format string
	Output string
}

// Option is a functional option for configuring the logger.
type Option func(*Config)

// WithLogLevel sets the log level.
func WithLogLevel(level string) Option {
	return func(cfg *Config) {
		cfg.Level = LogLevelFromString(level)
	}
}

// WithLogFormat sets the log format ("json" or "console").
func WithLogFormat(format string) Option {
	return func(cfg *Config) {
		cfg.Format = format
	}
}

// WithLogOutput sets the log output ("stdout", "stderr", or a file path).
func WithLogOutput(output string) Option {
	return func(cfg *Config) {
		cfg.Output = output
	}
}

// NewLogger creates a new logger instance with the given options.
func NewLogger(opts ...Option) Logger {
	// Default configuration
	cfg := &Config{
		Level:  zerolog.InfoLevel,
		Format: "json",
		Output: "stdout",
	}

	// Apply options
	for _, opt := range opts {
		opt(cfg)
	}

	// Set the global log level
	zerolog.SetGlobalLevel(cfg.Level)

	// Configure logger
	var logger zerolog.Logger
	if cfg.Format == "console" {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	} else {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	// Configure output
	if cfg.Output == "stderr" {
		logger = logger.Output(os.Stderr)
	} else if cfg.Output != "stdout" {
		file, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to open log file, using stdout")
		} else {
			logger = logger.Output(file)
		}
	}

	return &ZerologLogger{logger: logger, cfg: cfg}
}

// NewLoggerWithConfig creates a logger using the provided Config object.
func NewLoggerWithConfig(cfg *Config) Logger {
	return NewLogger(
		WithLogLevel(cfg.Level.String()), // Convert zerolog.Level to string
		WithLogFormat(cfg.Format),
		WithLogOutput(cfg.Output),
	)
}

// LogLevelFromString converts a string to a zerolog.Level.
func LogLevelFromString(level string) zerolog.Level {
	switch strings.ToLower(level) {
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
