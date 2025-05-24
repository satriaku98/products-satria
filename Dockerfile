FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .  

RUN go build -o myapp ./cmd

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/docs ./docs

EXPOSE 8080
CMD ["./myapp"]
