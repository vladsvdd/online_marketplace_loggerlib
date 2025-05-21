package online_marketplace_loggerlib

import "time"

type RequestContext struct {
	TraceID   string
	UserID    string
	StartedAt time.Time
}
