// Package logger logger/formate_error.go
package logger

import (
	"github.com/mdobak/go-xerrors"
	"log/slog"
	"path/filepath"
	"strings"
)

type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	if err, ok := a.Value.Any().(error); ok {
		a.Value = fmtErr(err)
	}
	return a
}

// returns a slog.Value with keys `msg` and `trace`. If the error
// does not implement interface { StackTrace() errors.StackTrace }, the `trace`
// key is omitted.
func fmtErr(err error) slog.Value {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("msg", err.Error()))

	// Проверяем, есть ли стек, если нет — оборачиваем
	if len(xerrors.StackTrace(err)) == 0 {
		err = xerrors.New(err)
	}

	frames := marshalStack(err)

	if frames != nil {
		groupValues = append(groupValues, slog.Any("trace", frames))
	}

	return slog.GroupValue(groupValues...)
}

// marshalStack extracts stack frames from the error
func marshalStack(err error) []stackFrame {
	trace := xerrors.StackTrace(err)

	if len(trace) == 0 {
		return nil
	}

	frames := trace.Frames()
	stackFrames := make([]stackFrame, 0, len(frames))

	for _, v := range frames {
		// Пропускаем системные кадры логгера. Начинаем записывать из места, где возникла ошибка
		source := filepath.Join(filepath.Base(filepath.Dir(v.File)), filepath.Base(v.File))
		if strings.Contains(source, "logger/logger.go") || strings.Contains(source, "slog/") {
			continue
		}

		stackFrames = append(stackFrames, stackFrame{
			Source: source,
			Func:   filepath.Base(v.Function),
			Line:   v.Line,
		})
	}

	return stackFrames
}
