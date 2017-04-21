@echo off

set curpath=%cd%
cd schema
init.cmd > initdb.log && cd %curpath%\tests && cd %curpath% && echo "Initdb OK"
