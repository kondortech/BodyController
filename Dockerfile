FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go build -o run ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/run /app/run

EXPOSE 8000

ENTRYPOINT [ "/app/run" ]