FROM golang:alpine3.22 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/main cmd/http/main.go

FROM ubuntu:22.04

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/main /app/bin/main 

RUN apt-get update && apt-get install -y libc6

RUN apt-get install -y netcat

EXPOSE 8080

# Use the runserver.sh script as the entry point
CMD ["/app/bin/main"]