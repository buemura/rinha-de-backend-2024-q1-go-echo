# Base
FROM golang:1.22-alpine3.19 as base

WORKDIR /app

COPY go.mod go.sum ./
COPY . . 
RUN go build -o main ./cmd/http

# Binary
FROM alpine:3.18 as binary
COPY --from=base /app/main .

CMD ["./main"]