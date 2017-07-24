@echo off

set curpath=%cd%
cd schema

init.cmd > initdb.log && ^
cd %curpath%\tests && ^
echo "Successful initdb for main-service" & ^
cd %curpath%
