package online_marketplace_loggerlib

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLogger_Infof(t *testing.T) {
	tmpFile := "./logs/test_log_info.log"

	log, err := NewLogger(
		WithFilePath(tmpFile),
		WithDebugMode(false),
		WithFormat(FormatJSON),
	)
	if err != nil {
		t.Fatalf("NewLogger failed: %v", err)
	}

	defer func() {
		err := log.Close()
		if err != nil {
			return
		}
		err = os.Remove(tmpFile)
		if err != nil {
			return
		}
	}()

	traceID := "test-trace"
	log.Info(traceID, "Test message: %s", "info")

	time.Sleep(50 * time.Millisecond) // дать время на запись

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Test message") {
		t.Error("Log file does not contain expected message")
	}
}

func TestLogger_Warningf(t *testing.T) {
	tmpFile := "./logs/test_log_warning.log"

	log, err := NewLoggerBuilder().WithFilePath(tmpFile).WithDebugMode(false).Build()
	if err != nil {
		t.Fatalf("NewLogger failed: %v", err)
	}

	defer func() {
		err := log.Close()
		if err != nil {
			return
		}
		err = os.Remove(tmpFile)
		if err != nil {
			return
		}
	}()

	traceID := "warn-trace"
	log.Warning(traceID, "Warning message: %s", "warn")

	time.Sleep(50 * time.Millisecond)

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Warning message") {
		t.Error("Log file does not contain expected warning message")
	}
}

func TestLogger_Errorf(t *testing.T) {
	tmpFile := "./logs/test_log_error.log"

	log, err := NewLogger(
		WithFilePath(tmpFile),
		WithDebugMode(false),
	)
	if err != nil {
		t.Fatalf("NewLogger failed: %v", err)
	}

	defer func() {
		err := log.Close()
		if err != nil {
			return
		}
		err = os.Remove(tmpFile)
		if err != nil {
			return
		}
	}()

	traceID := "error-trace"
	log.Error(traceID, "Error message: %s", "error")

	time.Sleep(50 * time.Millisecond)

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Error message") {
		t.Error("Log file does not contain expected error message")
	}
}

func TestLogger_Debugf(t *testing.T) {
	tmpFile := "./logs/test_log_debug.log"

	log, err := NewLogger(
		WithFilePath(tmpFile),
		WithDebugMode(true),
	)
	if err != nil {
		t.Fatalf("NewLogger failed: %v", err)
	}

	defer func() {
		err := log.Close()
		if err != nil {
			return
		}
		err = os.Remove(tmpFile)
		if err != nil {
			return
		}
	}()

	traceID := "debug-trace"
	log.Debug(traceID, "Debug message: %s", "debug")

	time.Sleep(50 * time.Millisecond)

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Debug message") {
		t.Error("Log file does not contain expected debug message")
	}
}

func TestLogger_With(t *testing.T) {
	tmpFile := "./logs/test_log_with.log"

	log, err := NewLogger(
		WithFilePath(tmpFile),
		WithDebugMode(false),
	)
	if err != nil {
		t.Fatalf("NewLogger failed: %v", err)
	}

	defer func() {
		err := log.Close()
		if err != nil {
			return
		}
		err = os.Remove(tmpFile)
		if err != nil {
			return
		}
	}()

	log2 := log.With("foo", "bar")
	log2.Info("trace", "message with context")

	time.Sleep(50 * time.Millisecond)

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "foo") {
		t.Error("Expected field 'foo' not found in log output")
	}
}

func TestLogger_WithContext(t *testing.T) {
	tmpFile := "./logs/test_log_ctx.log"

	log, err := NewLogger(
		WithFilePath(tmpFile),
		WithDebugMode(false),
	)
	if err != nil {
		t.Fatalf("NewLogger failed: %v", err)
	}

	defer func() {
		err := log.Close()
		if err != nil {
			return
		}
		err = os.Remove(tmpFile)
		if err != nil {
			return
		}
	}()

	ctx := context.Background()
	rc := &RequestContext{
		TraceID:   "ctx-trace",
		UserID:    "user-123",
		RequestID: "req-456",
		StartedAt: time.Now().Add(-1 * time.Second),
	}

	ctx = log.NewRequestContext(ctx, rc)
	log2 := log.WithContext(ctx)

	log2.Info(rc.TraceID, "contextual log")

	time.Sleep(50 * time.Millisecond)

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	data := string(content)
	if !strings.Contains(data, "ctx-trace") ||
		!strings.Contains(data, "user-123") ||
		!strings.Contains(data, "req-456") ||
		!strings.Contains(data, "startedAt") ||
		!strings.Contains(data, "duration") {
		t.Error("Expected context fields not found in log output")
	}
}
