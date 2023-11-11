FROM golang:1.21
SHELL ["/bin/bash", "-c"]

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p pkg
COPY pkg/ ./pkg

RUN mkdir -p services
COPY services/ ./services

