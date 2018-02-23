FROM golang:1.8

WORKDIR /go/src/data-service

COPY ./docker-conf ./conf
COPY ./controllers ./controllers 
COPY ./models ./models
COPY ./routers ./routers  
COPY ./main.go ./main.go  
COPY ./auth ./auth
COPY ./Gopkg.toml ./Gopkg.toml
COPY ./Gopkg.lock ./Gopkg.lock
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init
RUN dep ensure

RUN go install -v 

CMD ["data-service"]