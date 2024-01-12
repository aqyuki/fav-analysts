package logging

import (
	"context"
	"log/slog"
	"testing"
)

func Test_NewLogger(t *testing.T) {
	t.Run("NewLoggerFromEnv should return a logger in develop mode", func(t *testing.T) {
		t.Setenv(EnvLoggingMode, ModeDevelopment)
		actual := NewLoggerFromEnv()
		if actual == nil {
			t.Error("NewLoggerFromEnv should return a logger but received nil")
		}
	})

	t.Run("NewLoggerFromEnv should return a logger in production mode", func(t *testing.T) {
		t.Setenv(EnvLoggingMode, "")
		actual := NewLoggerFromEnv()
		if actual == nil {
			t.Error("NewLoggerFromEnv should return a logger but received nil")
		}
	})

	t.Run("NewLoggerWithConfig should return a logger", func(t *testing.T) {
		actual := NewLoggerWithConfig(true, "debug")
		if actual == nil {
			t.Error("NewLoggerWithConfig should return a logger but received nil")
		}
	})
}

func Test_convertLoggingLevel(t *testing.T) {
	t.Parallel()

	t.Run("convertLoggingLevel should return debug level", func(t *testing.T) {
		actual := covertLoggingLevel(LevelDebug)
		if actual != slog.LevelDebug {
			t.Errorf("convertLoggingLevel should return debug level but received %s", actual)
		}
	})

	t.Run("convertLoggingLevel should return info level", func(t *testing.T) {
		actual := covertLoggingLevel(LevelInfo)
		if actual != slog.LevelInfo {
			t.Errorf("convertLoggingLevel should return info level but received %s", actual)
		}
	})

	t.Run("convertLoggingLevel should return warn level", func(t *testing.T) {
		actual := covertLoggingLevel(LevelWarn)
		if actual != slog.LevelWarn {
			t.Errorf("convertLoggingLevel should return warn level but received %s", actual)
		}
	})

	t.Run("convertLoggingLevel should return error level", func(t *testing.T) {
		actual := covertLoggingLevel(LevelError)
		if actual != slog.LevelError {
			t.Errorf("convertLoggingLevel should return error level but received %s", actual)
		}
	})

	t.Run("convertLoggingLevel should return info level when unknown strings", func(t *testing.T) {
		actual := covertLoggingLevel("")
		if actual != slog.LevelInfo {
			t.Errorf("convertLoggingLevel should return info level but received %s", actual)
		}
	})
}

func Test_Context(t *testing.T) {
	t.Parallel()

	t.Run("Context should return a logger which is same as given logger", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := NewLoggerWithConfig(true, LevelDebug)

		actual := WithLogger(ctx, logger)
		if actual == nil {
			t.Errorf("WithLogger should return a context but received nil")
		}

		if UnwrapContext(actual) != logger {
			t.Errorf("UnwrapContext should return a logger which is same as given logger")
		}
	})

	t.Run("Context should returns a default logger", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		actual := UnwrapContext(ctx)
		if actual == nil {
			t.Errorf("UnwrapContext should return a logger but received nil")
		} else if actual != DefaultLogger() {
			t.Errorf("UnwrapContext should return a default logger")
		}
	})
}
