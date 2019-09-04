FROM golang:latest

WORKDIR /go/src/blog/

COPY ./package/nats ./src/cmd/nats
COPY ./src/cmd/cmdmanager/grpcConfig.yaml ./src/cmd/nats
