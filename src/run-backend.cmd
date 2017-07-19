@echo off
start "MAIN" /D main-service run-main-service.cmd
start "DATA" /D data-service run-data-service.cmd
start "AUTH" /D auth-service run-auth-service.cmd
start "FILE" /D file-service run-file-service.cmd