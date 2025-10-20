FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go build ./cmd/main.go


#multi-stage building...

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
