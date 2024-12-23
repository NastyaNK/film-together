FROM ubuntu:22.04 AS builder
LABEL authors="anastasia"


FROM golang:1.21.1-alpine AS builder

WORKDIR /root/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o myapp .

RUN chmod +x /root/myapp

FROM alpine:latest

COPY --from=builder /root/myapp /root/myapp

RUN adduser -D myuser
USER myuser

# Запускаем приложение
ENTRYPOINT ["/root/myapp"]


