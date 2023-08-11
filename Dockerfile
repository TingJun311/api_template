FROM golang:1.21-alpine

WORKDIR /app

# Copy the local go.mod and go.sum files to the container's working directory
COPY go.mod .
COPY go.sum .

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Golang application
RUN go build -o app

# Expose the port that your application will listen on
EXPOSE 8080:8080

# Command to run your application
CMD ["./app"]