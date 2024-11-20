package logging

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is an interface that wraps zerolog's basic logging methods and more.
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
	Fatalf(format string, v ...interface{})
	Printf(format string, v ...interface{})
	Sprintf(format string, v ...interface{}) string
}

// ZerologLogger implements the Logger interface using zerolog.
type ZerologLogger struct {
	logger zerolog.Logger
}

// NewZerologLogger creates a new instance of ZerologLogger.
func NewZerologLogger() Logger {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	return &ZerologLogger{
		logger: log.Output(zerolog.ConsoleWriter{Out: os.Stdout}),
	}
}

// Debug logs a message at the debug level.
func (l *ZerologLogger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// Info logs a message at the info level.
func (l *ZerologLogger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Warn logs a message at the warn level.
func (l *ZerologLogger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

// Error logs a message at the error level.
func (l *ZerologLogger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

// Fatal logs a message at the fatal level and exits the application.
func (l *ZerologLogger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

// Debugf logs a formatted message at the debug level.
func (l *ZerologLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debug().Msgf(format, v...)
}

// Infof logs a formatted message at the info level.
func (l *ZerologLogger) Infof(format string, v ...interface{}) {
	l.logger.Info().Msgf(format, v...)
}

// Warnf logs a formatted message at the warn level.
func (l *ZerologLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warn().Msgf(format, v...)
}

// Errorf logs a formatted message at the error level.
func (l *ZerologLogger) Errorf(format string, v ...interface{}) {
	l.logger.Error().Msgf(format, v...)
}

// Fatalf logs a formatted message at the fatal level and exits the application.
func (l *ZerologLogger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatal().Msgf(format, v...)
}

// Printf logs a formatted message at the info level.
func (l *ZerologLogger) Printf(format string, v ...interface{}) {
	l.logger.Info().Msgf(format, v...)
}

// Sprintf returns a formatted string using fmt.Sprintf
func (l *ZerologLogger) Sprintf(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}
