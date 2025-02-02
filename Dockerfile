# Start from the latest golang base image
FROM golang:1.22.5-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first; they are less frequently changed than source code, so Docker can cache this layer
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

FROM alpine:latest AS runner
RUN apk -U add yt-dlp && apk -U add ffmpeg

WORKDIR /app

COPY --from=builder /app/main .
# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]