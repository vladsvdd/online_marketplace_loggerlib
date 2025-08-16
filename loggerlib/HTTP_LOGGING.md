# HTTP Request Logging

Этот модуль предоставляет универсальную функцию для логирования HTTP-запросов к внешним API.

## Возможности

- **Унифицированное логирование** - единый формат для всех HTTP-запросов
- **Автоматическое построение URL** - параметры автоматически добавляются в URL
- **Гибкость** - возможность логирования любых API через одну функцию
- **Структурированность** - все данные логируются в удобном для анализа виде

## Структура HTTPRequestLog

```go
type HTTPRequestLog struct {
    Method      string                 `json:"method"`       // HTTP метод
    URL         string                 `json:"url"`          // Полный URL запроса
    Params      map[string]interface{} `json:"params,omitempty"`      // Параметры запроса
    RequestData map[string]interface{} `json:"request_data,omitempty"` // Данные запроса
    Response    interface{}            `json:"response,omitempty"`     // Ответ API
    StatusCode  int                    `json:"status_code,omitempty"`  // HTTP статус код
    Error       error                  `json:"error,omitempty"`        // Ошибка если есть
    Service     string                 `json:"service"`      // Название сервиса
}
```

## Использование

### Логирование VK API

```go
func (s *Service) logVKRequest(ctx context.Context, method string, params api.Params, response interface{}, err error) {
    // Конвертируем api.Params в map[string]interface{} для логирования
    paramsMap := make(map[string]interface{})
    for key, value := range params {
        paramsMap[key] = value
    }
    
    // Используем универсальную функцию логирования
    s.log.LogAPIRequestWithURL(ctx, "VK", method, "https://api.vk.com/method/", paramsMap, nil, response, 0, err)
}
```

### Логирование Telegram API

```go
func (s *Service) logTelegramRequest(ctx context.Context, method string, data map[string]interface{}, responseBody []byte, statusCode int, err error) {
    // Используем универсальную функцию логирования
    s.log.LogAPIRequestWithURL(ctx, "Telegram", method, "https://api.telegram.org/bot", nil, data, nil, statusCode, err)
}
```

### Логирование любого API

```go
func (s *Service) logCustomAPIRequest(ctx context.Context, method, baseURL string, params map[string]interface{}, response interface{}, err error) {
    s.log.LogAPIRequestWithURL(ctx, "CustomAPI", method, baseURL, params, nil, response, 0, err)
}
```

## Преимущества

1. **DRY принцип** - нет дублирования кода логирования
2. **Единообразие** - все логи выглядят одинаково
3. **Легкость поддержки** - изменения в логике логирования в одном месте
4. **Расширяемость** - легко добавить новые сервисы
5. **Структурированность** - логи легко парсить и анализировать

## Пример вывода

```json
{
  "level": "INFO",
  "msg": "[VK] API request details",
  "method": "photos.getAlbums",
  "url": "https://api.vk.com/method/photos.getAlbums?owner_id=-123&need_system=1",
  "service": "VK",
  "params": {
    "owner_id": -123,
    "need_system": 1
  },
  "response": {...},
  "error": null
}
```

## Добавление новых сервисов

Для добавления нового сервиса используйте универсальный метод `LogAPIRequestWithURL`:

```go
func (s *Service) logCustomServiceRequest(ctx context.Context, method string, data map[string]interface{}, response interface{}, err error) {
    s.log.LogAPIRequestWithURL(ctx, "CustomService", method, "https://api.customservice.com/", nil, data, response, 0, err)
}
```

## Параметры LogAPIRequestWithURL

```go
LogAPIRequestWithURL(
    ctx context.Context,                    // Контекст
    serviceName string,                     // Название сервиса (VK, Telegram, etc.)
    method string,                          // HTTP метод или endpoint
    baseURL string,                         // Базовый URL API
    params map[string]interface{},          // Параметры запроса (для GET)
    requestData map[string]interface{},     // Данные запроса (для POST)
    response interface{},                   // Ответ API
    statusCode int,                         // HTTP статус код
    err error,                              // Ошибка если есть
)
```
