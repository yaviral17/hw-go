# Use the official Golang image as the base image
FROM golang:1.22.2 as builder
# FROM postgres:12.1
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Copy wait-for-it script
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Build the Go app
RUN go build -o main .

# Expose port 6789 to the outside world
EXPOSE 6789

# Command to run the executable
CMD /wait-for-it.sh db:5432 -- ./main