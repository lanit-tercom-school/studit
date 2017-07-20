@echo off
title DATA
cd ..
cd ..
set GOPATH=%cd%
set PATH=%PATH%;%GOPATH%\bin
cd src/data-service
bee run -downdoc=true -gendoc=true