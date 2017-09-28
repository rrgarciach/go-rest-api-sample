FROM ntboes/golang-gin

ADD . /go/src/app

WORKDIR /go/src/app

RUN \
       go get -u github.com/gorilla/mux github.com/go-redis/redis

CMD ["gin", "run", "main.go"]
