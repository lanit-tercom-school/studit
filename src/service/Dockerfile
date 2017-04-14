FROM golang:1.8.1

COPY bin/service.* /usr/bin/backend
COPY src/service/conf/ /usr/bin/backend/conf/

EXPOSE 8081

CMD ["service"]