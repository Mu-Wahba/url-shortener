# Use the official Go image as a base
FROM golang:1.22-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./ 


# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from a smaller image
FROM alpine:latest  

# Install necessary dependencies to run the Go app
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/
COPY .env .

# Copy the Go app from the builder stage
COPY --from=builder /app/main .

# Expose the port the app will run on (assume 8080 in this case)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

