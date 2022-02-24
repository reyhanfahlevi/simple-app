FROM golang:1.14

RUN go get -v github.com/markbates/refresh

WORKDIR /go/src/github.com/reyhanfahlevi/simple-app
