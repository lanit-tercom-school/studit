FROM golang:1.8

WORKDIR /go/src/data-service

COPY ./docker-conf ./conf
COPY ./controllers ./controllers 
COPY ./models ./models
COPY ./routers ./routers  
COPY ./swagger ./swagger  
COPY ./main.go ./main.go  
COPY ./auth ./auth
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init
RUN dep ensure

RUN go install -v 

CMD ["data-service"]