FROM golang:1.24.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install dependencies first (better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy src code
COPY . .

# Build application
RUN go build -o /app/server ./cmd/server/main.go

# --- Final runtime image ---
FROM alpine:latest

WORKDIR /app

# Copy the binary and config
COPY --from=builder /app/server .
COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./server"]
