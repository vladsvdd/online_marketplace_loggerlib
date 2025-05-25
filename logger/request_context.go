// Package logger request_context.go
package logger

import "time"

type RequestContext struct {
	TraceID   string
	UserID    string
	RequestID string
	StartedAt time.Time
}
