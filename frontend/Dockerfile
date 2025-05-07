# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY frontend/ ./frontend/
COPY shared/ ./shared/
RUN CGO_ENABLED=0 GOOS=linux go build -o ./frontend/main ./frontend/main.go

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/frontend/public ./public
COPY --from=builder /app/frontend/views ./views
COPY --from=builder /app/frontend/main .
EXPOSE 8080
CMD ["./main"]
