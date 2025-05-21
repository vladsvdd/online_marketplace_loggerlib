package loggerlib

import "time"

type RequestContext struct {
	TraceID   string
	UserID    string
	StartedAt time.Time
}
