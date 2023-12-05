#!/bin/sh

set -e

host="$1"
shift

until PGPASSWORD=pa55word psql -h "$host" -U "greenlight" -c '\q'; do
  >&2 echo "Postgres is unavailable"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec "$@"