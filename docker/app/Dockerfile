# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy .env config from docker application directory to the working Directory inside the container
COPY ./docker/app/.env .env

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go web
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/app ./cmd/web
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/migrate ./cmd/migrate

# Start a new stage from scratch
FROM alpine:latest
RUN apk update && apk --no-cache add ca-certificates
RUN apk add --no-cache bash postgresql-client

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/bin/app /app/
COPY --from=builder /app/bin/migrate /app/
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/.env /app/

COPY --from=builder /app/docker/app/bin /usr/local/bin/app/
RUN chmod +x /usr/local/bin/app/*

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["/app/app"]