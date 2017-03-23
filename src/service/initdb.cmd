@echo off
set curpath=%cd%
cd schema
init.cmd > initdb.log & cd %curpath%\tests & go test > init_tests.log & cd %curpath%
