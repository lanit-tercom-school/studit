export GOPATH=$PWD/../../..
export PATH=$PATH:$GOPATH/bin
export PGPASSWORD=postgres
export PGUSER=postgres

dropdb -e --host=localhost --port=5432 --username=$PGUSER --no-password --if-exists studit
createdb -e --owner=$PGUSER --host=localhost --port=5432 --username=$PGUSER --no-password studit
psql -a --host=localhost --port=5432 --username=$PGUSER --no-password --file=data_base_init.sql studit

echo Database studit was correctly initialized

go run db_init_data.go