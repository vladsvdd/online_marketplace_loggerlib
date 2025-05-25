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

func (h *traceHandler) Enabled(_ context.Context, level slog.Level) bool {
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

type LogFormat string

const (
	FormatJSON LogFormat = "json"
	FormatText LogFormat = "text"
)

type Config struct {
	FilePath string
	IsDebug  bool
	Format   LogFormat
}

// Option Вариативные параметры через ...Option (Go-идиоматично)
// Это гибкий и расширяемый способ, который часто используется в библиотеках.
type Option func(*Config)

func WithFilePath(path string) Option {
	return func(cfg *Config) {
		cfg.FilePath = path
	}
}

func WithDebugMode(debug bool) Option {
	return func(cfg *Config) {
		cfg.IsDebug = debug
	}
}

func WithFormat(fmt LogFormat) Option {
	return func(cfg *Config) {
		cfg.Format = fmt
	}
}

func MakeLogger(opts ...Option) (*Logger, error) {
	cfg := Config{
		FilePath: LogFilePath,
		IsDebug:  false,
		Format:   FormatJSON,
	}
	for _, o := range opts {
		o(&cfg)
	}
	return buildLogger(cfg)
}

// Builder pattern
type Builder struct {
	cfg Config
}

func NewLoggerBuilder() *Builder {
	return &Builder{
		cfg: Config{
			FilePath: LogFilePath,
			IsDebug:  false,
			Format:   FormatJSON,
		},
	}
}

func (b *Builder) WithFilePath(path string) *Builder {
	b.cfg.FilePath = path
	return b
}

func (b *Builder) WithDebugMode(debug bool) *Builder {
	b.cfg.IsDebug = debug
	return b
}

func (b *Builder) WithFormat(fmt LogFormat) *Builder {
	b.cfg.Format = fmt
	return b
}

func (b *Builder) Build() (*Logger, error) {
	return buildLogger(b.cfg)
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
