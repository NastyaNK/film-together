FROM golang:1.21.1-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o myapp .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/myapp .
COPY --from=builder /app/resources /app/resources

RUN adduser -D myuser
USER myuser

ENTRYPOINT ["./myapp"]
