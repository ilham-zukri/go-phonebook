#!/bin/bash

# ----- CONFIG -----
DRIVER="postgres"
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="phonebook_db"
MIGRATIONS_DIR="./database/migrations"

# ----- ASK FOR CREDENTIALS -----
read -p "DB username: " DB_USER
read -s -p "DB password: " DB_PASS
echo ""

DB_URL="${DRIVER}://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# ---- COMMAND HANDLER ----

case "$1" in
    create)
        if [ -z "$2" ]; then
            echo "❌ Migration name required"
            echo "Usage: ./migrate.sh create migration_name"
            exit 1
        fi

        migrate create -ext sql -dir "${MIGRATIONS_DIR}" -seq "$2"
        ;;
    
    up)
        migrate -path "${MIGRATIONS_DIR}" -database "$DB_URL" up
        ;;
    
    down)
        migrate -path "${MIGRATIONS_DIR}" -database "$DB_URL" down
        ;;
    
    *)
        echo "Usage:"
        echo "  ./migrate.sh create migration_name"
        echo "  ./migrate.sh up"
        echo "  ./migrate.sh down"
        ;;

esac
