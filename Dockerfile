# Stage 1: Build the Go application
FROM golang:1.22-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o loanengine .

# Stage 2: Create a small image for running the Go application
FROM alpine:3.15

# Install ca-certificates and tzdata for time zone support
RUN apk add --no-cache ca-certificates tzdata

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=build /app/loanengine .

# Create log directory and set permissions
RUN mkdir -p /var/log/loanEngine && chmod -R 777 /var/log/loanEngine

# Set the time zone to Asia/Kolkata
ENV TZ=Asia/Kolkata

# Expose port 5002 to the outside world
EXPOSE 5002

# Command to run the executable
CMD ["./loanengine"]
