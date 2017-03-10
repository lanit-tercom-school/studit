cd ..
cd ..
cd ..
set GOPATH=%cd%
set PATH=%PATH%;%GOPATH%\bin;C:\Program Files\PostgreSQL\9.6\bin
cd src/service/schema
set PGPASSWORD=postgres
set PGUSER=postgres
dropdb -e --host=localhost --port=5432 --username=postgres --no-password --if-exists studit
createdb -e --owner=postgres --host=localhost --port=5432 --username=postgres --no-password studit
psql -a --host=localhost --port=5432 --username=postgres --no-password --file=data_base_init.sql studit
echo TABLE studit was correctly initialized
go run db_init_data.go