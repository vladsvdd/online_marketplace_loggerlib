// Package online_marketplace_libs logger/types.go
package online_marketplace_loggerlib

import (
	"log/slog"
	"os"
	"time"
)

type LogFormat string

const (
	FormatJSON LogFormat = "json"
	FormatText LogFormat = "text"
)

const DefaultLogFilePath = "./logs/log.log"

type RequestContext struct {
	TraceID   string
	RequestID string
	UserID    string
	Method    string
	Path      string
	Status    int
	Ip        string
	StartedAt time.Time
}

type Config struct {
	FilePath string
	IsDebug  bool
	Format   LogFormat
}

type Logger struct {
	l    *slog.Logger
	file *os.File
}

type traceHandler struct {
	slog.Handler
	level slog.Level
}
