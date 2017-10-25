#!/bin/sh
cd ../
export GOPATH=`pwd`
export PATH=${PATH}:${GOPATH}/bin
cd src/file-service
bee run -downdoc=true -gendoc=true
