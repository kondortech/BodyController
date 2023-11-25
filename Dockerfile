FROM golang:1.21

WORKDIR /

COPY ./ .

RUN go mod download