# -------------------------------
# Stage 1: Builder
# -------------------------------
FROM golang:1.24.5 AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник
RUN go build -o main ./cmd/main.go

# -------------------------------
# Stage 2: Runtime
# -------------------------------
FROM debian:bookworm-slim

# Рабочая директория
WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/main .

# Копируем конфиг и .env внутрь контейнера
COPY configs /root/configs
COPY .env /root/.env


ENV CONFIG_FILE=/root/configs/config

# Запуск приложения
CMD ["./main"]
