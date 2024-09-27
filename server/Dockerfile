FROM golang:1.18

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/main.go

# Use a smaller base image for the final stage
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=0 /app/myapp .

# Uncomment bellow line if local Docker deployment in PaaS autosets 
# EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]