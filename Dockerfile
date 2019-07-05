FROM golang:latest

ENV GO111MODULE=on

WORKDIR /go/src/github.com/masashi1014/go_sample
COPY . .

RUN go get github.com/oxequa/realize
RUN go build