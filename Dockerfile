# Заготовка для Dockerfile, будет доработана позже
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /go_url_shortener

EXPOSE 8080

CMD ["/go_url_shortener"]
