@echo off
start "MAIN" /D main-service run_studit_back.cmd
start "AUTH" /D auth-service run_studit_auth.cmd
