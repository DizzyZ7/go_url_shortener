# Используем официальный образ Golang в качестве базового
FROM golang:1.18-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для загрузки зависимостей
COPY go.mod .
COPY go.sum .

# Загружаем Go-модули
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Собираем исполняемый файл
RUN go build -o go_url_shortener .

# Открываем порт 8080 для Gin-сервера
EXPOSE 8080

# Запускаем приложение
CMD ["./go_url_shortener"]
