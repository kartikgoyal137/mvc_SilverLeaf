#!/bin/sh
set -e

DB_URL="mysql://${DBUSER}:${DBPASS}@tcp(${DBHOST}:3306)/${DBNAME}?parseTime=true"

echo "Waiting for database to be ready..."
sleep 2

echo "Running database migrations..."

/migrate -database "$DB_URL" -path /migrations up

echo "Migrations complete. Starting server..."

exec "$@"
