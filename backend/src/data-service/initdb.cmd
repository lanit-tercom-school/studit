@echo off

set curpath=%cd%
cd schema
init.cmd > initdb.log && echo "Successful initdb for data-service"