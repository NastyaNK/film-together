#!/bin/bash

if [ -z "$1" ]; then
    echo "Использование: ./migrations.sh [up|down]"
    exit 1
fi

DB_URL="postgres://anastasia:2553@localhost:5433/postgres?sslmode=disable"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
MIGRATIONS_PATH="$SCRIPT_DIR/pkg/repository/migrations"


case "$1" in
    up)
        echo "Запуск миграций..."
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" up
        ;;
    down)
        echo "Откатываем миграции..."
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" down -all
        ;;
    *)
        echo "Неверный аргумент! Используйте './start.sh up' или './start.sh down'"
        exit 1
        ;;
esac
