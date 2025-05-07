# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/frontend ./frontend/main.go

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/frontend .
COPY --from=builder /app/frontend/public ./public
COPY --from=builder /app/frontend/views ./views

EXPOSE 8080
CMD ["./frontend"]