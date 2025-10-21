# Sử dụng image Go để build ứng dụng
FROM golang:1.24-alpine AS builder


WORKDIR /app

# Chỉ copy file module trước để cache dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ mã nguồn vào container
COPY . .

# Build app thành file thực thi
RUN go build -o main .

# Image chạy app
FROM alpine:latest
WORKDIR /app

# Copy file binary từ stage builder
COPY --from=builder /app/main .
COPY --from=builder /app/logs ./logs

EXPOSE 8080

CMD ["./main"]
