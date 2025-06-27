# Stage 1 — Build the Go app
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather-backend main.go

# Stage 2 — Run in a minimal image
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary and assets
COPY --from=builder /app/weather-backend .
COPY --from=builder /app/internal/config/config.go ./internal/config/
COPY --from=builder /app/internal/util ./internal/util

# Copy any needed files (e.g., empty DB file if required)
# NOTE: SQLite will auto-create the DB file if missing
VOLUME ["/app/data"]

# Expose the server port
EXPOSE 8080

CMD ["./weather-backend"]
