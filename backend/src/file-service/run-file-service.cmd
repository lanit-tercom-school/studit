@echo off
title FILE
cd ..
cd ..
set GOPATH=%cd%
set PATH=%PATH%;%GOPATH%\bin
cd src/file-service
bee run -downdoc=true -gendoc=true