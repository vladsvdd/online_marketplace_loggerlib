// Package logger logger/logger.go
package logger

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

const LogFilePath = "./logs/log.log"

type Logger struct {
	l    *slog.Logger
	file *os.File
}

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

func (s *Logger) Infof(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Info(message, args...)
}

func (s *Logger) Warningf(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Warn(message, args...)
}

func (s *Logger) Errorf(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Error(message, args...)
}

func (s *Logger) Debugf(traceID, message string, args ...interface{}) {
	s.l.With("traceID", traceID).Debug(message, args...)
}

type traceHandler struct {
	slog.Handler
	level slog.Level
}

func (h *traceHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *traceHandler) Handle(ctx context.Context, r slog.Record) error {
	logID := uuid.NewString()
	r.AddAttrs(slog.String("logID", logID))
	return h.Handler.Handle(ctx, r)
}

func (s *Logger) Close() error {
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}

func getLevel(l slog.Leveler) slog.Level {
	if level, ok := l.(slog.Level); ok {
		return level
	}
	return l.Level()
}

func MakeLogger(filePath string, isDebug bool) (*Logger, error) {
	if filePath == "" {
		filePath = LogFilePath
	}

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать директорию логов: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл логов: %v", err)
	}

	opts := &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
		Level:       slog.LevelInfo,
	}
	if isDebug {
		opts.Level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(file, opts)

	th := &traceHandler{
		Handler: handler,
		level:   getLevel(opts.Level),
	}

	return &Logger{
		l:    slog.New(th),
		file: file,
	}, nil
}
