# Stage 1: Build the Go app inside an official Go image
FROM golang:1.24.2 AS builder

WORKDIR /app

# Install build tools for CGO
RUN apt-get update && apt-get install -y gcc

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the full source code
COPY . .

# Build the app using CGO (required by go-sqlite3)
RUN CGO_ENABLED=1 GOOS=linux go build -o weather-backend main.go

# Stage 2: Run using Debian (compatible with glibc)
FROM debian:bullseye-slim

WORKDIR /app

# Install runtime dependencies for the binary
RUN apt-get update && apt-get install -y libc6 ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary from build stage
COPY --from=builder /app/weather-backend .

# Optional: create volume for SQLite persistence
VOLUME ["/app/data"]

# Expose your app port
EXPOSE 8080

# Run the binary
CMD ["./weather-backend"]
