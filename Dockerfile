# Builder Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Runtime Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations ./migrations

# Install postgresql-client for migration check if needed, but not strictly required by app logic
# However, we need to ensure the migration files are reachable.

EXPOSE 8080

CMD ["./main"]
