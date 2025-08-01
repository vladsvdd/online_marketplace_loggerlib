# Автоматическое и ручное версионирование

Этот репозиторий содержит workflows для автоматического и ручного управления версиями.

## Доступные Workflows

### `auto-version.yml` - Автоматическое версионирование
- Запускается при каждом push в ветку `master`
- Автоматически увеличивает patch версию при изменениях кода
- Игнорирует изменения только в документации

### `tag-latest.yml` - Ручное версионирование
- Запускается вручную через GitHub Actions
- Позволяет выбрать тип версии (major/minor/patch)
- Для создания релизов и специальных версий

## Автоматическое версионирование

### Триггер
- Запускается при каждом `push` в ветку `master`
- Игнорирует изменения только в документации (`.md`, `README`, `.gitignore`)

### Логика
1. **Проверка изменений**: Анализирует, есть ли изменения в коде
2. **Проверка флага**: Ищет `[skip version]` в сообщении коммита
3. **Версионирование**: Если нужно, увеличивает patch версию (1.0.0 → 1.0.1)

## Использование автоматического версионирования

**Примечание**: Автоматическое версионирование всегда использует `patch` версию (1.0.0 → 1.0.1).

### Обычный коммит (создаст новую patch версию)
```bash
git add .
git commit -m "feat: add new logging functionality"
git push origin master
# Результат: v1.0.0 → v1.0.1
```

### Коммит без версионирования
```bash
git commit -m "docs: update readme [skip version]"
git push origin master
# Результат: версия не изменится
```

### Изменения только в документации (версия не изменится)
```bash
git commit -m "docs: update installation guide"
git push origin master
# Результат: версия не изменится (только .md файлы)
```

### Примеры автоматического patch версионирования:
```bash
# Исправление бага
git commit -m "fix: resolve memory leak in logger"
git push origin master
# v1.0.0 → v1.0.1

# Новая функция
git commit -m "feat: add OAuth2 authentication"
git push origin master
# v1.0.1 → v1.0.2

# Рефакторинг
git commit -m "refactor: improve error handling"
git push origin master
# v1.0.2 → v1.0.3
```


git add .
git commit -m "docs: add comprehensive versioning documentation"

# 2. Создать минорную версию
./scripts/version.sh minor
# Результат: v0.5.5 → v0.6.0

# 3. Запушить коммит
git push origin master