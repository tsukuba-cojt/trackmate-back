# -- Build stage ---
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o migration ./migrations/migration.go
RUN go build -o api-server .

# --- Run stage ---
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/migration .
COPY --from=builder /app/api-server .

# マイグレーション実行後にAPIサーバー起動
CMD ["sh", "-c", "./migration && ./api-server"]