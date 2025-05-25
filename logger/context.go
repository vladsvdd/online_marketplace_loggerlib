// Package logger logger/context.go
package logger

import (
	"context"
)

type ctxKeyType string

const ctxKey ctxKeyType = "request_context_key"

func (s *Logger) NewRequestContext(parent context.Context, rc *RequestContext) context.Context {
	return context.WithValue(parent, ctxKey, rc)
}

func (s *Logger) GetRequestContext(ctx context.Context) *RequestContext {
	val := ctx.Value(ctxKey)
	if rc, ok := val.(*RequestContext); ok {
		return rc
	}
	return nil
}
