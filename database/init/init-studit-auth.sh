dropdb -e --host=localhost --port=5432 --username="$POSTGRES_USER" --no-password --if-exists studit_auth
createdb -e --owner="$POSTGRES_USER" --host=localhost --port=5432 --username="$POSTGRES_USER" --no-password studit_auth
