FROM golang:latest

WORKDIR /go/src/blog/

COPY ./package/cmdmanager ./src/cmd/cmdmanager
COPY ./src/cmd/cmdmanager/grpcConfig.yaml ./src/cmd/cmdmanager
COPY ./src/cmd/api/config.local.yaml ./src/cmd/cmdmanager

EXPOSE 50051