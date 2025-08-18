#!/bin/sh
set -e

DB_URL="mysql://${DBUSER}:${DBPASS}@tcp(${DBHOST}:3306)/${DBNAME}?parseTime=true"

echo "Waiting for database to be ready at ${DBHOST}:3306..."

until mysql -h "${DBHOST}" -u "${DBUSER}" -p"${DBPASS}" --skip-ssl -e 'SELECT 1'; do
  >&2 echo "Database is unavailable - sleeping"
  sleep 1
done

>&2 echo "Database is up - continuing..."

echo "Running database migrations..."
/migrate -database "$DB_URL" -path /migrations up

echo "Migrations complete. Starting server..."
exec "$@"