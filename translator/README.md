# Translator Package

Пакет для работы с переводами в приложении online_marketplace.

## Использование

### Базовое использование

```go
import "github.com/vladsvdd/online_marketplace_libs/translator"

// Получение перевода
message := translator.Translate(translator.KeyTitleRequired, translator.RU)
// Результат: "Заголовок обязателен"
```

### Использование констант (рекомендуется)

Вместо строковых литералов рекомендуется использовать константы для избежания опечаток:

```go
import "github.com/vladsvdd/online_marketplace_libs/translator"

// ✅ Правильно - используем константы
message := translator.Translate(translator.KeyTitleRequired, translator.RU)
error := errors.New(translator.Translate(translator.KeyItemCreationFailed, translator.RU))

// ❌ Неправильно - строковые литералы могут содержать опечатки
message := translator.Translate("title_required", translator.RU)
error := errors.New(translator.Translate(translator.KeyItemCreationFailed, translator.RU))
```

### Доступные константы

#### Универсальные поля
```go
translator.KeyTitleRequired           // "title_required"
translator.KeyDescriptionRequired     // "description_required"
translator.KeyLocationRequired        // "location_required"
translator.KeyCategoryRequired        // "category_required"
translator.KeyPriceRequired           // "price_required"
translator.KeySalaryRequired          // "salary_required"
translator.KeyContactInfoRequired     // "contact_info_required"
translator.KeyPropertyTypeRequired    // "property_type_required"
translator.KeyRoomCountRequired       // "room_count_required"
translator.KeyFloorRequired           // "floor_required"
translator.KeyTotalFloorsRequired     // "total_floors_required"
translator.KeyAreaRequired            // "area_required"
translator.KeyExperienceRequired      // "experience_required"
translator.KeyEducationRequired       // "education_required"
translator.KeySkillsRequired          // "skills_required"
translator.KeyCoverLetterRequired     // "cover_letter_required"
translator.KeyMessageRequired         // "message_required"
translator.KeyNotificationTypeRequired = "notification_type_required"
translator.KeyNotificationEnabledRequired = "notification_enabled_required"
translator.KeySearchQueryRequired     // "search_query_required"
```

#### Универсальные ID
```go
translator.KeyIDRequired             // "id_required"
translator.KeyUserIDRequired         // "user_id_required"
translator.KeyGroupIDRequired        // "group_id_required"
translator.KeyLocationIDRequired     // "location_id_required"
translator.KeySiteCategoryIDRequired // "site_category_id_required"
```

#### Универсальные операции
```go
translator.KeyItemCreatedSuccess     // "item_created_success"
translator.KeyItemUpdatedSuccess     // "item_updated_success"
translator.KeyItemDeletedSuccess     // "item_deleted_success"
translator.KeyItemNotFound           // "item_not_found"
translator.KeyItemCreationError      // "item_creation_error"
translator.KeyItemUpdateError        // "item_update_error"
translator.KeyItemDeleteError        // "item_delete_error"
translator.KeyItemCreationFailed     // "item_creation_failed"
translator.KeyItemUpdateFailed       // "item_update_failed"
translator.KeyItemDeleteFailed       // "item_delete_failed"
```

#### Аутентификация
```go
translator.KeyRegistrationError              // "registration_error"
translator.KeyRegistrationCompletedSuccessfully // "registration_completed_successfully"
translator.KeyUserNotFound                   // "user_not_found"
translator.KeyEmailNotConfirmed              // "email_not_confirmed"
translator.KeyInvalidLoginOrPassword         // "invalid_login_or_password"
translator.KeyTokenGenerationError           // "token_generation_error"
translator.KeyRefreshTokenInvalid            // "refresh_token_invalid"
translator.KeyRefreshTokenConvertError       // "refresh_token_convert_error"
translator.KeyJWTTokenCreationError          // "jwt_token_creation_error"
translator.KeyUserAlreadyExists              // "user_already_exists"
translator.KeyInvalidSigningMethod           // "invalid_signing_method"
translator.KeyTokenClaimsError               // "token_claims_error"
```

### Примеры использования в коде

#### Валидация данных
```go
func validateVacancy(data models.Vacancy) error {
    switch {
    case data.Title == nil || *data.Title == "":
        return errors.New(translator.Translate(translator.KeyTitleRequired, translator.RU))
    case data.Description == nil || *data.Description == "":
        return errors.New(translator.Translate(translator.KeyDescriptionRequired, translator.RU))
    case data.LocationId == nil || *data.LocationId == "":
        return errors.New(translator.Translate(translator.KeyLocationIDRequired, translator.RU))
    }
    return nil
}
```

#### Обработка ошибок
```go
func createVacancy(data models.Vacancy) (*models.Vacancy, error) {
    if err := validateVacancy(data); err != nil {
        return nil, err
    }
    
    result, err := repository.Create(data)
    if err != nil {
        return nil, fmt.Errorf(translator.Translate(translator.KeyItemCreationFailed, translator.RU))
    }
    
    return result, nil
}
```

#### HTTP ответы
```go
func (h *Handler) CreateVacancy(c *gin.Context) {
    var data models.Vacancy
    if err := c.ShouldBindJSON(&data); err != nil {
        handler_helper.RespondWithError(c, http.StatusBadRequest, 
            translator.Translate(translator.KeyTitleRequired, lang))
        return
    }
    
    result, err := h.service.Create(data)
    if err != nil {
        handler_helper.RespondWithError(c, http.StatusInternalServerError, 
            translator.Translate(translator.KeyItemCreationFailed, lang))
        return
    }
    
    c.JSON(http.StatusCreated, result)
}
```

### Преимущества использования констант

1. **Безопасность типов** - компилятор проверит правильность ключей
2. **Автодополнение** - IDE предложит доступные ключи
3. **Рефакторинг** - легко переименовать ключи без поиска по всему коду
4. **Предотвращение опечаток** - невозможно случайно написать неправильный ключ
5. **Документация** - константы служат документацией доступных переводов

### Миграция с строковых литералов

Для постепенной миграции существующего кода:

1. Замените строковые литералы на константы по одному файлу
2. Используйте поиск и замену: `"title_required"` → `translator.KeyTitleRequired`
3. Проверьте компиляцию после каждого изменения
4. Обновите тесты, если они используют строковые литералы

### Добавление новых переводов

При добавлении новых переводов:

1. Добавьте константу в блок `const` в `translator.go`
2. Добавьте перевод в карту `translations`
3. Используйте константу в коде

```go
// В translator.go
const (
    KeyNewTranslation = "new_translation"
)

var translations = map[string]map[Language]string{
    KeyNewTranslation: {
        EN: "New translation",
        RU: "Новый перевод",
    },
}
``` 