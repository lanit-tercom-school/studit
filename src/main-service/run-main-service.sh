#!/bin/sh
cd ../
export GOPATH=`pwd`
export PATH=${PATH}:${GOPATH}/bin
cd src/main-service
go run main.go
