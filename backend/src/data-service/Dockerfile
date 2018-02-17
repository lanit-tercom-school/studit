FROM golang:1.8

WORKDIR /go/src/data-service

COPY . .
COPY ./docker-conf ./conf
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init
RUN dep ensure

RUN go install -v 

CMD ["data-service"]