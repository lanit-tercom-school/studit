FROM golang:1.8

WORKDIR /go/src/auth-service

COPY ./docker-conf ./conf
COPY ./controllers ./controllers 
COPY ./models ./models
COPY ./routers ./routers  
COPY ./main.go ./main.go  
COPY ./Gopkg.toml ./Gopkg.toml
COPY ./Gopkg.lock ./Gopkg.lock
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/beego/bee
RUN dep ensure

CMD bee run