// Package online_marketplace_loggerlib request_context.go
package logger

import "time"

type RequestContext struct {
	TraceID   string
	UserID    string
	RequestID string
	StartedAt time.Time
}
