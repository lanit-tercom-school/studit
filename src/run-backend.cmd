@echo off
title RUN BACK
start /d main-service run-main-service.cmd
start /d data-service run-data-service.cmd
start /d auth-service run-auth-service.cmd
start /d file-service run-file-service.cmd