# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY go.mod go.sum ./
RUN go mod tidy
COPY backend/ /app/
WORKDIR /app/backend/cmd
RUN go build -o todo-backend .

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/backend/cmd/todo-backend /root/
EXPOSE 8080
CMD ["./todo-backend"]
