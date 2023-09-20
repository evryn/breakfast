# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.20-alpine3.18 as builder

ARG BUILD_VERSION
ARG BUILD_DATE
ARG BUILD_REF

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./... \
 && go build -o /go/bin/app -v .

# A layer for running automated tests
FROM builder as test

RUN go test ./...

# A layer with final executable
FROM alpine:3.18

RUN apk --no-cache add ca-certificates \
 && adduser -S -u 1000 -s /bin/bash -h /home/version-forge version-forge

WORKDIR /home/version-forge/

COPY . /home/version-forge/
COPY --from=builder /go/bin/app /home/version-forge/version-forge

USER version-forge
EXPOSE 8080
ENTRYPOINT ["/home/version-forge/version-forge"]
