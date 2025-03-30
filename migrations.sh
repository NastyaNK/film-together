#!/bin/bash

if [ -z "$1" ]; then
    echo "Использование: ./start.sh [up|down]"
    exit 1
fi

DB_URL="postgres://anastasia:2553@localhost:5432/postgres?sslmode=disable"
MIGRATIONS_PATH="migrations"

case "$1" in
    up)
        echo "Запуск миграций..."
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" up
        ;;
    down)
        echo "Откатываем миграции..."
        migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" down
        ;;
    *)
        echo "Неверный аргумент! Используйте './start.sh up' или './start.sh down'"
        exit 1
        ;;
esac
