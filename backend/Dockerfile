FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go build ./cmd/main.go

CMD ["./main"]
