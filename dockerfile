# Use an official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY src/ .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run your application
CMD ["go", "run", "main.go"]
