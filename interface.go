package loggerlib

import (
	"context"
)

type Logger interface {
	NewRequestContext(parent context.Context, requestContext *RequestContext) context.Context
	GetRequestContext(ctx context.Context) *RequestContext
	With(args ...any) Logger
	WithContext(ctx context.Context) Logger
	Infof(traceID, message string, args ...interface{})
	Warningf(traceID, message string, args ...interface{})
	Errorf(traceID, message string, args ...interface{})
	Debugf(traceID, message string, args ...interface{})
}
