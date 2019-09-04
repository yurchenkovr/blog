FROM golang:1.12-alpine

LABEL maintainer="Vitalii Yurchenko <vitalik.yurchenko2000@gmail.com>"

RUN apk add --no-cache \
        libc6-compat

WORKDIR /go/src/blog/

COPY ./package/api ./src/cmd/api
COPY ./src/cmd/api/entrypoint.sh ./src/cmd/api
COPY ./src/cmd/cmdmanager/grpcConfig.yaml ./src/cmd/api

EXPOSE 8080
