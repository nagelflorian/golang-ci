FROM golang:1.8.1

ADD . /go/src/github.com/nagelflorian/golang-ci

RUN go install github.com/nagelflorian/golang-ci

CMD ["/go/bin/golang-ci"]

EXPOSE 3000
