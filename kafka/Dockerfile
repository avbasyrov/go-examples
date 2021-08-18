# Сборка приложения
FROM golang:latest as builder
RUN mkdir /build
WORKDIR /build
COPY . .
RUN go build -a -o kafkaapp ./main.go
ENTRYPOINT [ "/build/kafkaapp" ]
