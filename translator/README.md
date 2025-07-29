# Translator

Модуль для мультиязычных переводов (локализации) фраз.

## Как использовать

1. **Загрузите переводы из файла:**

```go
import "online_marketplace_libs/translator"

err := translator.LoadTranslations("/path/to/translations.json")
if err != nil {
    // обработка ошибки
}
```

2. **Получите перевод по ключу и языку:**

```go
msg := translator.Translate("user_not_found", translator.RU) // "Пользователь не найден"
msg2 := translator.Translate("user_not_found", translator.EN) // "User not found"
```

## Формат файла переводов (JSON)

```json
{
  "user_not_found": {
    "en": "User not found",
    "ru": "Пользователь не найден"
  },
  "invalid_password": {
    "en": "Invalid password",
    "ru": "Неверный пароль"
  }
}
```

## Добавление новых языков и фраз
- Просто добавьте новые ключи и переводы в `translations.json`.
- Для поддержки других языков используйте их ISO-коды (например, "de", "fr"). 