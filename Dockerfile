FROM alpine:latest

WORKDIR /app

COPY ./dist/myapp .
COPY ./resources .
COPY ./config ./config

RUN adduser -D myuser
USER myuser

EXPOSE 8080
ENTRYPOINT ["./myapp"]