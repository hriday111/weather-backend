# Stage 1: build
FROM golang:1.24.2-bookworm AS builder

WORKDIR /app
RUN apt-get update && apt-get install -y gcc

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o weather-backend main.go

# Stage 2: run
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y libc6 ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/weather-backend .

EXPOSE 8080
CMD ["./weather-backend"]
