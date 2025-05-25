Реализация логера

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


Создание логгера с поддержкой подходов — через опции (`Option`), и через билдер (`LoggerBuilder`)
Оба подхода создают один и тот же объект `Logger`, используя общую функцию `buildLogger`.
### Через опции:
```
log, err := MakeLogger(
    WithFilePath("./tmp/log.json"),
    WithDebugMode(true),
)
```

### Через билдер:
```
log, err := NewLoggerBuilder().
    WithFilePath("./tmp/log.json").
    WithDebugMode(true).
    Build()
```