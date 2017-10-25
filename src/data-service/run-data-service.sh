#!/bin/sh
cd ../
export GOPATH=`pwd`
export PATH=${PATH}:${GOPATH}/bin
cd src/data-service
bee run -downdoc=true -gendoc=true
