FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 7000
CMD ["sh", "-c", "go run main.go $(cat ${ENV_FILE})"]
ENV ENV_FILE .env.prod