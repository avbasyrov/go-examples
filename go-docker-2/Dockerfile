# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /my-app

FROM alpine:latest as run

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /my-app .
# EXPOSE 8080 # раскомментировать если хотим опубликовать какой-либо порт

CMD [ "./my-app" ]

