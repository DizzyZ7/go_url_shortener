# Используем многоступенчатую сборку для уменьшения размера итогового образа

# Сборка
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go_url_shortener .

# Финальный образ
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go_url_shortener .
COPY --from=builder /app/views ./views

EXPOSE 8080

CMD ["./go_url_shortener"]
