# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY backend/ /app/
RUN go build -o todo-backend .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/todo-backend /root/
EXPOSE 8080
CMD ["./todo-backend"]
