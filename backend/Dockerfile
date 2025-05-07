# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY backend/ ./backend/
COPY shared/ ./shared/
RUN CGO_ENABLED=0 GOOS=linux go build -o ./backend/cmd/main ./backend/cmd/main.go

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/backend/main .
EXPOSE 8080
CMD ["./main"]
