export GOPATH=$PWD/../..
export PATH=$PATH:$GOPATH/bin

go get github.com/astaxie/beego           && \
go get github.com/beego/bee               && \
go get github.com/lib/pq                  && \
go get github.com/smartystreets/goconvey  && \
go get github.com/robbert229/jwt          && \
echo "Packeges were installed successfully"