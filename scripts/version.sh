#!/bin/bash

# Скрипт для автоматического управления версиями
# Использование: ./version.sh [major|minor|patch]

set -euo pipefail  # Более строгий режим выполнения

# Проверка правильного remote URL (поддерживает токены аутентификации)
REMOTE_URL=$(git remote get-url origin)
if [[ ! "$REMOTE_URL" =~ (online_marketplace_libs)(\.git)?$ ]]; then
    echo "Ошибка: неправильный remote URL. Ожидается online_marketplace_libs[.git]" >&2
    echo "Текущий URL: $REMOTE_URL" >&2
    echo "Исправьте командой: git remote set-url origin https://github.com/vladsvdd/online_marketplace_libs" >&2
    exit 1
fi

# Функция для получения текущей версии
get_current_version() {
    # Ищем только теги версий (vX.Y.Z), исключая другие теги
    git tag --sort=-v:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | head -n 1 || echo "v0.0.0"
}

# Функция для увеличения версии
increment_version() {
    local version=$1
    local increment_type=$2

    # Проверяем формат версии
    if [[ ! "$version" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        echo "Ошибка: некорректный формат версии '$version'" >&2
        exit 1
    fi

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

# Проверяем, что нет незакоммиченных изменений
if ! git diff-index --quiet HEAD --; then
    echo "Ошибка: есть незакоммиченные изменения" >&2
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

# Проверяем, существует ли уже такой тег
if git show-ref --tags --quiet --verify "refs/tags/$new_version"; then
    echo "Ошибка: тег $new_version уже существует" >&2
    exit 1
fi

# Создаем тег с аннотацией
git tag -a "$new_version" -m "Version $new_version"
git tag -f latest

# Пушим теги
echo "Отправка тегов на удаленный репозиторий..."
git push origin "$new_version"
git push -f origin latest

echo "Успешно:"
echo " - Создан тег версии: $new_version"
echo " - Обновлен тег latest"