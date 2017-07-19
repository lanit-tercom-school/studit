@echo off
title MAIN
cd ..
cd ..
set GOPATH=%cd%
set PATH=%PATH%;%GOPATH%\bin
cd src/main-service
bee run -downdoc=true -gendoc=true