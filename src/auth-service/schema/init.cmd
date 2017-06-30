@echo off

set postgres64=C:\Program Files\PostgreSQL\9.6\bin
set postgres32=C:\Program Files (x86)\PostgreSQL\9.6\bin

IF EXIST "%postgres64%" set path=%path%;%postgres64%
IF EXIST "%postgres32%" set path=%path%;%postgres32%

set GOPATH=%cd%/../../..
set PATH=%PATH%;%GOPATH%\bin
set PGPASSWORD=postgres
set PGUSER=postgres
dropdb -e --host=localhost --port=5432 --username=postgres --no-password --if-exists studit_auth && echo "dropdb OK" ^
&& createdb -e --owner=postgres --host=localhost --port=5432 --username=postgres --no-password studit_auth && echo "createdb OK" ^
&& psql --host=localhost --port=5432 --username=postgres --no-password --file=auth-schema.sql studit_auth && echo "psql OK" ^
&& go run auth-schema.go && echo "All is OK"
