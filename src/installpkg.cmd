@echo off
set GOPATH=%cd%/..
go get github.com/astaxie/beego           && ^
go get github.com/beego/bee               && ^
go get github.com/lib/pq                  && ^
go get github.com/smartystreets/goconvey  && ^
go get github.com/robbert229/jwt          && ^
go get github.com/google/uuid             && ^
go get -u github.com/nleof/goyesql        && ^
echo "Packeges were installed successfully"