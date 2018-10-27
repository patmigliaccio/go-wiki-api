FROM golang:1.3

ENV PORT 80

ADD . /go/src/github.com/patmigliaccio/untitled
WORKDIR /go/src/github.com/patmigliaccio/untitled

RUN go get github.com/go-martini/martini
RUN go get github.com/martini-contrib/render
RUN go get gopkg.in/mgo.v2
RUN go get github.com/martini-contrib/binding

EXPOSE 80
CMD go run server.go