#!/usr/bin/env bash

RETRIES=5

host="$1"
shift
cmd="$@"

until psql -h $host -U $DB_USER -p $DB_PASSWORD -d DB_NAME -c "select 1" > /dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
 echo "Waiting for postgres server, $((RETRIES--)) remaining attempts..."
 sleep 1
done

>&2 echo "postgres is up - executing command"
exec $cmd