export GOPATH=$PWD/../../..
export PATH=$PATH:$GOPATH/bin
export PGPASSWORD=postgres
export PGUSER=postgres

dropdb -e --host=localhost --port=5432 --username=postgres --no-password --if-exists studit_file && echo "dropdb OK" \
&& createdb -e --owner=postgres --host=localhost --port=5432 --username=postgres --no-password studit_file && echo "createdb OK" \
&& psql --host=localhost --port=5432 --username=postgres --no-password --file=data_base_init.sql studit_file && echo "psql OK" \
&& go run db_init_data.go && echo "Default data inserted" && echo "All is OK"
