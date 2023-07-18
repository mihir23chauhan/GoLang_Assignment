FROM golang:1.20

WORKDIR /app

COPY . .

# Build the Go application inside the container
RUN go build -o main .

# Expose the port that the REST API will listen on
EXPOSE 4000

# Command to run the Go application
CMD ["./main"]