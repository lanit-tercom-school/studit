@echo off

set curpath=%cd%
cd schema
init.cmd > initdb.log && cd %curpath%\tests && echo "Successful initdb for file-service" & cd %curpath%
