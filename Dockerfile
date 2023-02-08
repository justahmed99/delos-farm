# Use an official Go image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the entire project directory into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port that the Go app will listen on
EXPOSE 9090

# Specify the command to run when the container starts
CMD ["./main"]