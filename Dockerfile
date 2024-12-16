# Stage 1: Build the Go application
FROM golang:1.21 AS builder
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application and build
COPY . .
RUN go build -o app .

# Stage 2: Create the runtime image
FROM ubuntu:22.04

# Install ffmpeg and dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates


# Copy the built Go binary from the builder stage
COPY --from=builder /app/app /usr/local/bin/app


# Expose port 8080 for the Go application
EXPOSE 8080

# Command to run the Go application
CMD ["app"]



