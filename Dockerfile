FROM golang:1.8.3-alpine3.6

ADD . /home

WORKDIR /home

RUN \
       apk add --no-cache bash git openssh && \
       go get -u github.com/minio/minio-go github.com/go-redis/redis



CMD ["go","run","app.go"]
