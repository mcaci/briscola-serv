FROM golang:1.19.0-alpine

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN cp /go/bin/dlv /dlv
COPY . /project