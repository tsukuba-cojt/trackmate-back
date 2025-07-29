# -- Build stage ---
FROM golang:1.24.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o api-server .
RUN go build -o migration ./migrations/migration.go

# --- Run stage ---
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/api-server .
COPY --from=builder /app/migration .

# マイグレーション実行後にAPIサーバー起動
CMD ["sh", "-c", "./migration && ./api-server"]