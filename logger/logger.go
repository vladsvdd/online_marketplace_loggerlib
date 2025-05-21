package logger

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vladsvdd/online_marketplace_loggerlib"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

const LogFilePath = "./logs/log.log"

type Logger struct {
	l *slog.Logger
}

// Implement loggerlib.Logger
func (s *Logger) With(args ...any) online_marketplace_loggerlib.Logger {
	return &Logger{
		l: s.l.With(args...),
	}
}

func (s *Logger) WithContext(ctx context.Context) online_marketplace_loggerlib.Logger {
	rc := s.GetRequestContext(ctx)
	if rc == nil {
		return s
	}

	attrs := []any{
		"traceID", rc.TraceID,
		"userID", rc.UserID,
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

type TraceHandler struct {
	slog.Handler
}

func (h *TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	logID := uuid.NewString()
	r.AddAttrs(slog.String("logID", logID))
	return h.Handler.Handle(ctx, r)
}

func MakeLogger(filePath string) (online_marketplace_loggerlib.Logger, error) {
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

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{ReplaceAttr: replaceAttr})
	traceHandler := &TraceHandler{Handler: handler}

	return &Logger{
		l: slog.New(traceHandler),
	}, nil
}
