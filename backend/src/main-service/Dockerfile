FROM golang:1.8

WORKDIR /go/src/main-service

COPY ./docker-conf ./conf
COPY ./handler ./handler 
COPY ./helpers ./helpers
COPY ./mutations ./mutations  
COPY ./objects ./objects  
COPY ./queries ./queries  
COPY ./root ./root  
COPY ./main.go ./main.go  
COPY ./Gopkg.toml ./Gopkg.toml
COPY ./Gopkg.lock ./Gopkg.lock
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init
RUN dep ensure

RUN go install -v 

CMD ["main-service"]