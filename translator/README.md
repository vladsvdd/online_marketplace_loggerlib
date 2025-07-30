# Translator

Простая библиотека для перевода текста на разные языки.

## Использование

```go
import "online_marketplace_libs/translator"

// Получить перевод
message := translator.Translate("user_not_found", translator.RU)
// Результат: "Пользователь не найден"

// Если перевод не найден, возвращается ключ
message := translator.Translate("unknown_key", translator.RU)
// Результат: "unknown_key"
```

## Поддерживаемые языки

- `translator.EN` - Английский
- `translator.RU` - Русский (по умолчанию)

## Доступные ключи переводов

Библиотека содержит переводы для всех основных сообщений системы:

- Ошибки валидации (`input_data_incorrect_format`, `username_required`, etc.)
- Сообщения аутентификации (`user_not_found`, `invalid_login_or_password`, etc.)
- Сообщения токенов (`token_generation_error`, `refresh_token_invalid`, etc.)
- Общие ошибки (`database_error`, `validation_error`, etc.)
- Успешные операции (`item_created_success`, `favorite_added_success`, etc.)

## Особенности

- Переводы хранятся в памяти (не требуется загрузка из файла)
- Потокобезопасность через `sync.RWMutex`
- Fallback на русский язык, если перевод для запрошенного языка не найден
- Если ключ не найден, возвращается сам ключ 