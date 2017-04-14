FROM golang:1.8.1

COPY src/service/service /usr/bin/backend/
COPY src/service/conf/ /usr/bin/backend/conf/

EXPOSE 8081

CMD ["/usr/bin/service"]