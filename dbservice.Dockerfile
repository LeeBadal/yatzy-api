# Use the official Golang image as the base image
FROM golang:1.21.0-alpine3.18

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

RUN mkdir -p /app/migrations

# Copy the entire project directory to the container
COPY . .

COPY ./cmd/dbservice/migrations/ /app/migrations/

# Build the application binary for dbservice
RUN go build -o yatzy-dbservice ./cmd/dbservice/

# Expose port 50051 for the API
EXPOSE 50051

# Start the dbservice when the container starts
CMD ["./yatzy-dbservice"]
