// Package online_marketplace_loggerlib logger/logger.go
package online_marketplace_loggerlib

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func (s *Logger) With(args ...any) *Logger {
	return &Logger{
		l: s.l.With(args...),
	}
}

func (s *Logger) WithContext(ctx context.Context) *Logger {
	rc := s.GetRequestContext(ctx)
	if rc == nil {
		return s
	}

	attrs := []any{
		"userID", rc.UserID,
		"requestId", rc.RequestID,
	}

	if rc.TraceID != "" {
		attrs = append(attrs, "traceID", rc.TraceID)
	}

	if !rc.StartedAt.IsZero() {
		duration := time.Since(rc.StartedAt)
		attrs = append(attrs,
			"startedAt", rc.StartedAt,
			"duration_human", fmt.Sprintf("%.6f", duration.Seconds()),
			"duration", duration,
		)
	}

	return s.With(attrs...)
}

func (s *Logger) Info(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Info(message, args...)
}

func (s *Logger) Warning(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Warn(message, args...)
}

func (s *Logger) Error(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Error(message, args...)
}

func (s *Logger) Debug(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Debug(message, args...)
}

func (s *Logger) Close() error {
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}

func buildLogger(cfg Config) (*Logger, error) {
	err := os.MkdirAll(filepath.Dir(cfg.FilePath), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать директорию логов: %v", err)
	}

	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл логов: %v", err)
	}

	handlerOpts := &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
		Level:       slog.LevelInfo,
	}

	if cfg.IsDebug {
		handlerOpts.Level = slog.LevelDebug
	}

	var baseHandler slog.Handler
	switch cfg.Format {
	case FormatText:
		baseHandler = slog.NewTextHandler(file, handlerOpts)
	case FormatJSON:
		fallthrough
	default:
		baseHandler = slog.NewJSONHandler(file, handlerOpts)
	}

	th := &traceHandler{
		Handler: baseHandler,
		level:   handlerOpts.Level.Level(),
	}

	return &Logger{
		l:    slog.New(th),
		file: file,
	}, nil
}

func NewLogger(opts ...Option) (*Logger, error) {
	cfg := Config{
		FilePath: DefaultLogFilePath,
		IsDebug:  false,
		Format:   FormatJSON,
	}

	for _, o := range opts {
		o(&cfg)
	}

	return buildLogger(cfg)
}
