# Указываем базовый образ Go
FROM golang:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . ./
COPY .env ./

# Собираем бинарник
RUN go build -o /build

# Открываем порт
# EXPOSE 9090
EXPOSE 8080

# Запускаем приложение
CMD [ "/build" ]
