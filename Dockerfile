# Стадия сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o todo-app main.go


FROM alpine:latest

WORKDIR /app


COPY --from=builder /app/todo-app ./

EXPOSE 8080

CMD ["./todo-app"]
