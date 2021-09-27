# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /my-app

# EXPOSE 8080 # раскомментировать если хотим опубликовать какой-либо порт

CMD [ "/my-app" ]

