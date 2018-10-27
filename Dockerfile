FROM golang:1.7

ENV PORT 80

ADD . /go/src/github.com/patmigliaccio/go-wiki-api
WORKDIR /go/src/github.com/patmigliaccio/go-wiki-api

RUN go get github.com/gorilla/mux
RUN go get github.com/patmigliaccio/go-wikimedia

EXPOSE 80
CMD go run *.go