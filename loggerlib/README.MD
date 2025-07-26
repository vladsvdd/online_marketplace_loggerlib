Реализация логера

types.go - Основные типы и константы

options.go - Функциональные опции

builder.go - Builder паттерн

context.go - Работа с контекстом

handler.go - Кастомный обработчик логов

formate_error.go - Форматирование ошибок

logger.go - Основная реализация логгера

Example interface
```
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
```

### Пример использования:
```
logger, err := logger.MakeLogger(
	logger.WithFilePath("./custom.log"),
	logger.WithDebugMode(true),
)
if err != nil {
	panic(err)
}
defer logger.Close()

ctx := logger.NewRequestContext(context.Background(), &logger.RequestContext{
	TraceID:   "trace-123",
	RequestID: "req-456",
	UserID:    "user-789",
	StartedAt: time.Now(),
})

logger.Infof(ctx, "Processing request", "order_id", 123)
```