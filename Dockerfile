# Use an official Go runtime as the base image
FROM golang:1.21.3

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Build the Go application
RUN go build -o basic-tasks-api

# Expose the port your application listens on
EXPOSE 8080

# Run the Go application
CMD ["./basic-tasks-api"]
