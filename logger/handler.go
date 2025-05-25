// logger/handler.go
package logger

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (h *traceHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *traceHandler) Handle(ctx context.Context, r slog.Record) error {
	logID := uuid.NewString()
	r.AddAttrs(slog.String("logID", logID))
	return h.Handler.Handle(ctx, r)
}
