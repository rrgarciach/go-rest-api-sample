FROM ntboes/golang-gin

ADD ./src /go/src/app

WORKDIR /go/src/app

RUN sed -i 's/archive.ubuntu.com/mirrors.rit.edu/' /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get -y install xsltproc git make libxml2-utils openssh-client

RUN go get -u github.com/gorilla/mux github.com/go-redis/redis

CMD ["gin", "run", "main.go"]
