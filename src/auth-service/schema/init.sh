export GOPATH=$PWD/../../..
export PATH=$PATH:$GOPATH/bin
export PGPASSWORD=postgres
export PGUSER=postgres

dropdb -e --host=localhost --port=5432 --username=postgres --no-password --if-exists studit_auth && echo "dropdb OK" \
&& createdb -e --owner=postgres --host=localhost --port=5432 --username=postgres --no-password studit_auth && echo "createdb OK" \
&& psql --host=localhost --port=5432 --username=postgres --no-password --file=auth-schema.sql studit_auth && echo "psql OK" \
&& go run auth-schema.go && echo "Default data inserted" && echo "All is OK"
