@echo off
title AUTH
cd ..
cd ..
set GOPATH=%cd%
set PATH=%PATH%;%GOPATH%\bin
cd src/auth-service
bee run -downdoc=true -gendoc=true