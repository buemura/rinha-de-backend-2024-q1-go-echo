# Base
FROM golang:1.21-alpine3.18 as base

ARG PORT
ARG DATABASE_URL

WORKDIR /app

RUN echo "PORT=${PORT}" > .env
RUN echo "DATABASE_URL=${DATABASE_URL}" >> .env

COPY go.mod go.sum ./
COPY . . 
RUN go build -o main ./cmd/http

# Binary
FROM alpine:3.18 as binary
COPY --from=base /app/main .
COPY --from=base /app/.env .

CMD ["./main"]