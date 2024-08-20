## Stage 1: Build Stage

# Use the latest golang base image for the builder
FROM golang:alpine AS builder

# Set up the working directory
RUN mkdir -p /go/src/service
ADD . /go/src/service
WORKDIR /go/src/service

# Install dependencies
RUN apk add git
RUN go mod tidy

# Build Swagger docs (uncomment if needed)
RUN go install -v github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g main.go -o Library/Swagger/docs

# Run tests (uncomment if needed)
# RUN CGO_ENABLED=0 GOOS=linux go test ./Controller/...

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o main main.go
RUN chmod 755 /go/src/service/main

## Stage 2: Run Stage

# Use the latest alpine base image
FROM alpine:latest
LABEL maintainer="Putra Agastiya <ageztya.putra@gmail.com>"

# Create a non-root user
RUN addgroup -S putra && adduser -S putra -G putra

# Install necessary packages
RUN apk update && apk upgrade && apk add --no-cache tzdata

# Set environment variable from build argument
ARG TAG
ENV environment=$TAG

# Set up the application directory
RUN mkdir -p /app
WORKDIR /app

# Copy necessary files from the builder stage
COPY --chown=putra:putra --from=builder /go/src/service/environment /go/src/service/environment
COPY --chown=putra:putra --from=builder /go/src/service/main /app

# Switch to the non-root user
USER putra

# Expose port 7000 to the outside world
EXPOSE 7000

# Command to run the application
CMD /app/main -tag=$environment
