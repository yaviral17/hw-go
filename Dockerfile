# Use the official Golang image as the base image
FROM golang:1.22.2 as builder
# FROM postgres:12.1
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run the Go application
FROM golang:1.22.2

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 6789 to the outside world
EXPOSE 6789

# Command to run the executable
CMD ["./main"]