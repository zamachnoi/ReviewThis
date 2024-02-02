# Use an official Go runtime as a parent image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app for a linux target
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/viewthis-backend ./cmd/api

# Use a Docker multi-stage build to create a thin image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /go/bin/viewthis-backend .
COPY --from=builder /go/src/app/.env .

# Expose port 8080 to the outside world
EXPOSE 3001

# Command to run the executable
CMD ["./viewthis-backend"]
