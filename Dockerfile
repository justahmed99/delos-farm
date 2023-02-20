# # Use an official Go image as the base image
# FROM golang:latest

# # Set the working directory to /app
# WORKDIR /app

# # Copy the entire project directory into the container
# COPY . .

# # Build the Go app
# # RUN go build -o main .

# RUN go get github.com/joho/godotenv

# CMD ["sh", "-c", "go run main.go $(cat ${ENV_FILE})"]

# # CMD ["go", "run", "main.go", "--env-file", ".env-docker"]

# # Expose the port that the Go app will listen on
# EXPOSE 7000

# # Specify the command to run when the container starts
# # CMD ["./main"]

# Use an official Go image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the entire project directory into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port that the Go app will listen on
EXPOSE 7000

# Specify the command to run when the container starts
CMD ["sh", "-c", "go run main.go $(cat ${ENV_FILE})"]

# Set the default environment variable to use the .env file in production
ENV ENV_FILE .env.prod
