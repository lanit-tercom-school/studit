export GOPATH=$PWD/../..
export PATH=$PATH:$GOPATH/bin
export curpath=$PWD

cd schema
./init.sh > initdb.log & cd $curpath\tests & go test > init_tests.log & cd $curpath
