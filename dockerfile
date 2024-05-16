FROM golang:1.21.4-alpine3.18 AS builder

WORKDIR /app

# Copy only the necessary files for module downloading
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire project into the container
COPY . .

WORKDIR /app/app

# Build the Go application
RUN go build -o /app/app/main

# Final stage
FROM scratch

ARG APP_PORT
ENV APP_PORT=80

EXPOSE 80

WORKDIR /app/app

# Copy the binary from the builder stage
COPY --from=builder /app/app/main .

# Copy .env file into the container
COPY ../.env .

CMD ["/app/app/main"]
