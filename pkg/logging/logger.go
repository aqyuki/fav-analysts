package logging

import (
	"log/slog"
	"os"
	"strings"
)

const (
	// EnvLoggingMode is a environment variable name for logging mode
	EnvLoggingMode = "LOG_MODE"

	// ModeDevelopment is a development mode
	// Default mode is production mode
	ModeDevelopment = "develop"

	// EnvLoggingLevel is a environment variable name for logging level
	// If development mode is enabled, use debug level
	// Default level is info level
	EnvLoggingLevel = "LOG_LEVEL"

	// LevelDebug is a debug level
	LevelDebug = "debug"
	// LevelInfo is a info level
	LevelInfo = "info"
	// LevelWarn is a warn level
	LevelWarn = "warn"
	// LevelError is a error level
	LevelError = "error"
	// DefaultLevel is a default level
	DefaultLevel = LevelInfo
)

// NewLoggerFromEnv creates a new logger from configuration from environment variables.
func NewLoggerFromEnv() *slog.Logger {
	develop := strings.ToLower(strings.TrimSpace(os.Getenv(EnvLoggingMode))) == ModeDevelopment
	level := strings.ToLower(strings.TrimSpace(os.Getenv(EnvLoggingLevel)))
	return NewLoggerWithConfig(develop, level)
}

// NewLoggerWithConfig creates a new logger with config.
func NewLoggerWithConfig(develop bool, level string) *slog.Logger {
	var handler slog.Handler

	if develop {
		handler = slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdin, &slog.HandlerOptions{
			AddSource: false,
			Level:     covertLoggingLevel(level),
		})
	}
	return slog.New(handler)
}

// covertLoggingLevel converts string to slog.Level
func covertLoggingLevel(s string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
