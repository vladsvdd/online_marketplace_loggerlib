// Package logger logger/context.go
package logger

import (
	"context"
	"github.com/vladsvdd/online_marketplace_loggerlib"
)

type ctxKeyType string

const ctxKey ctxKeyType = "request_context_key"

func (s *Logger) NewRequestContext(parent context.Context, rc *online_marketplace_loggerlib.RequestContext) context.Context {
	return context.WithValue(parent, ctxKey, rc)
}

func (s *Logger) GetRequestContext(ctx context.Context) *online_marketplace_loggerlib.RequestContext {
	val := ctx.Value(ctxKey)
	if rc, ok := val.(*online_marketplace_loggerlib.RequestContext); ok {
		return rc
	}
	return nil
}
