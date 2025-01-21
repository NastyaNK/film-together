FROM golang:1.21.1-alpine AS builder //создание образа на основе линукса и версии языка

WORKDIR /app // папка в который все сохарним
COPY . . // путь файла который переносим  и путь куда переносим

RUN go build -o myapp . // строим и называем образ

FROM alpine:latest // создаем еще образ чтобы тут хранить все

WORKDIR /app // работает в этой папке
COPY --from=builder /app/myapp . // первый образ переносим в новый образ
COPY --from=builder /app/resources /app/resources // и ресурсы тоже

RUN adduser -D myuser //
USER myuser //

ENTRYPOINT ["./myapp"] //
