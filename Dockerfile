# Builder stage
FROM golang:1.20.3 AS builder
# Set the working directory
WORKDIR /app
# Copy go.mod and go.sum to workdir
COPY go.mod go.sum ./
# Install go modules inside the image
RUN go mod download
# Copy the source code
COPY app/ ./app
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o goapp app/infrastructure/main.go

# Production stage
FROM alpine:3.2 AS production
# Set the working directory
WORKDIR /app
# Copy the binary application from builder stage
COPY --from=builder /app/goapp ./
# Expose port
EXPOSE 8086
# Start application
CMD ["./goapp"]