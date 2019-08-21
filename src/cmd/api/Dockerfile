# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="Vitalii Yurchenko <vitalik.yurchenko2000@gmail.com>"

WORKDIR /go/src/blog/

# Copy go mod and sum files
#COPY ./go.mod ./go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

COPY ./vendor ./vendor
COPY ./src/cmd/migration ./src/cmd/migration
COPY ./src/cmd/api ./src/cmd/api
COPY ./src/infrastructure ./src/infrastructure
COPY ./src/models ./src/models
COPY ./src/repository ./src/repository
COPY ./src/usecases ./src/usecases

RUN go build -o ./src/cmd/migration/migration ./src/cmd/migration
RUN go build -o ./src/cmd/api/api ./src/cmd/api
# Build the Go app
# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
#CMD ["./src/cmd/migration/migration"]
#CMD ["./src/cmd/api/api"]
