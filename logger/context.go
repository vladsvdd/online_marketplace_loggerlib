package logger

import (
	"context"
	"loggerlib"
)

type ctxKeyType string

const ctxKey ctxKeyType = "request_context_key"

func (s *Logger) NewRequestContext(parent context.Context, rc *loggerlib.RequestContext) context.Context {
	return context.WithValue(parent, ctxKey, rc)
}

func (s *Logger) GetRequestContext(ctx context.Context) *loggerlib.RequestContext {
	val := ctx.Value(ctxKey)
	if rc, ok := val.(*loggerlib.RequestContext); ok {
		return rc
	}
	return nil
}
