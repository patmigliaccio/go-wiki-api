FROM golang:1.7

ENV PORT 80

ADD . /go/src/github.com/patmigliaccio/untitled
WORKDIR /go/src/github.com/patmigliaccio/untitled

RUN go get github.com/gorilla/mux
RUN go get github.com/patrickmn/go-wikimedia

EXPOSE 80
CMD go run *.go