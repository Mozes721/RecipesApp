# Use the official Go image as the base image
FROM golang:1.18

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application code
COPY . .

RUN go build -o main .

CMD ["./main"]

