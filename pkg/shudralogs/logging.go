package shudralogs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Logger defines the logging interface for the library.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	ErrorfNew(format string, v ...interface{}) error
	Fatalf(format string, v ...interface{})
	Printf(format string, v ...interface{})
	Sprintf(format string, v ...interface{}) string
	PrettyPrint(v interface{})
	SetLogLevelUnsafe(lvl string)
}

// ZerologLogger is a concrete implementation of Logger using zerolog.
type ZerologLogger struct {
	logger zerolog.Logger
	cfg    *Config
}

// Debug logs a debug message.
func (l *ZerologLogger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// Info logs an info message.
func (l *ZerologLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Warn logs a warning message.
func (l *ZerologLogger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

// Error logs an error message.
func (l *ZerologLogger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

// Fatal logs a fatal message and exits the application.
func (l *ZerologLogger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

// Debugf logs a formatted debug message.
func (l *ZerologLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debug().Msgf(format, v...)
}

// Infof logs a formatted info message.
func (l *ZerologLogger) Infof(format string, v ...interface{}) {
	l.logger.Info().Msgf(format, v...)
}

// Warnf logs a formatted warning message.
func (l *ZerologLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warn().Msgf(format, v...)
}

// Errorf logs a formatted error message.
func (l *ZerologLogger) Errorf(format string, v ...interface{}) {
	l.logger.Error().Msgf(format, v...)
}

// ErrorfNew returns an error created with a formatted message.
func (l *ZerologLogger) ErrorfNew(format string, v ...interface{}) error {
	return errors.New(l.Sprintf(format, v...))
}

// Fatalf logs a formatted fatal message and exits the application.
func (l *ZerologLogger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatal().Msgf(format, v...)
}

// Printf logs a formatted message at the info level.
func (l *ZerologLogger) Printf(format string, v ...interface{}) {
	l.logger.Info().Msgf(format, v...)
}

// Sprintf returns a formatted string.
func (l *ZerologLogger) Sprintf(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}

// PrettyPrint logs a formatted JSON representation of a value.
func (l *ZerologLogger) PrettyPrint(v interface{}) {
	prettyData, _ := json.MarshalIndent(v, "", "  ")
	l.Debug(string(prettyData))
}

// SetLogLevelUnsafe sets the log level dynamically at runtime.
func (l *ZerologLogger) SetLogLevelUnsafe(lvl string) {
	zerolog.SetGlobalLevel(LogLevelFromString(lvl))
}

// UpdateOptions allows updating logger settings dynamically by applying new options.
func (l *ZerologLogger) UpdateOptions(opts ...Option) {
	// Apply new options to the current configuration
	for _, opt := range opts {
		opt(l.cfg)
	}

	// Update the logger based on the new configuration
	var logger zerolog.Logger
	if l.cfg.Format == "console" {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	} else {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	// Update the output
	switch l.cfg.Output {
	case "stderr":
		logger = logger.Output(os.Stderr)
	default:
		if l.cfg.Output != "stdout" {
			file, err := os.OpenFile(l.cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				logger = logger.Output(file)
			} else {
				logger.Error().Err(err).Msg("Failed to open log file, using stdout")
			}
		}
	}

	// Update the log level
	zerolog.SetGlobalLevel(l.cfg.Level)

	// Assign the updated logger
	l.logger = logger
}
