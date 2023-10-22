# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.21-alpine3.18 as builder

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
 && adduser -S -u 1000 -s /bin/bash -h /home/breakfast breakfast

WORKDIR /home/breakfast/

COPY . /home/breakfast/
COPY --from=builder /go/bin/app /home/breakfast/breakfast

USER breakfast
EXPOSE 8080
ENTRYPOINT ["/home/breakfast/breakfast"]
