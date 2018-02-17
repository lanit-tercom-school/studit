#!/bin/bash
export GOPATH="$PWD/../../.."
export PATH="$PATH:$GOPATH/bin"
export PGPASSWORD=postgres
export PGUSER=postgres

dropdb -e --host=localhost --port=5432 --username=postgres --no-password --if-exists studit && echo "dropdb OK" \
&& createdb -e --owner=postgres --host=localhost --port=5432 --username=postgres --no-password studit && echo "createdb OK" \
&& psql --host=localhost --port=5432 --username=postgres --no-password --file=data_base_init.sql studit && echo "psql OK" \
&& go run db_init_data.go && echo "All is OK"
