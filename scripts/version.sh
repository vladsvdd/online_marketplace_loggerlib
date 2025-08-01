#!/bin/bash

# Скрипт для автоматического управления версиями
# Использование: ./version.sh [major|minor|patch]

set -e

# Функция для получения текущей версии
get_current_version() {
    # Ищем только теги версий (vX.Y.Z), исключая 'latest' и другие неверсионные теги
    git tag --sort=-v:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | head -n 1 || echo "v0.0.0"
}

# Функция для увеличения версии
increment_version() {
    local version=$1
    local increment_type=$2

    # Убираем префикс v
    version=${version#v}

    # Разбиваем на части
    IFS='.' read -r major minor patch <<< "$version"

    case $increment_type in
        "major")
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        "minor")
            minor=$((minor + 1))
            patch=0
            ;;
        "patch")
            patch=$((patch + 1))
            ;;
        *)
            echo "Неверный тип версии. Используйте: major, minor, patch" >&2
            exit 1
            ;;
    esac

    echo "v${major}.${minor}.${patch}"
}

# Проверяем, что находимся в git-репозитории
if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo "Ошибка: не найден git-репозиторий" >&2
    exit 1
fi

# Основная логика
if [ $# -eq 0 ]; then
    echo "Использование: $0 [major|minor|patch]" >&2
    echo "  major - увеличивает мажорную версию (1.0.0 -> 2.0.0)" >&2
    echo "  minor - увеличивает минорную версию (1.0.0 -> 1.1.0)" >&2
    echo "  patch - увеличивает патч версию (1.0.0 -> 1.0.1)" >&2
    exit 1
fi

increment_type=$1
current_version=$(get_current_version)
new_version=$(increment_version "$current_version" "$increment_type")

echo "Текущая версия: $current_version"
echo "Новая версия: $new_version"

# Создаем тег с аннотацией
git tag -a "$new_version" -m "Version $new_version"
git tag -f latest

# Пушим теги
git push origin "$new_version"
git push -f origin latest

echo "Успешно:"
echo " - Создан тег версии: $new_version"
echo " - Обновлен тег latest"