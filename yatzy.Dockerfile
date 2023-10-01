# Use the official Golang image as the base image
FROM golang:1.21.0-alpine3.18

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./


# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the application binary
RUN go build -o yatzy ./cmd/api-gateway/

# Expose port 8080 for the API
EXPOSE 8080

# Start the API when the container starts
CMD ["./yatzy"]