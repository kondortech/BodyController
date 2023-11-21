FROM golang:1.21

WORKDIR /

COPY Makefile ./

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p pkg
COPY ./pkg/ ./pkg

RUN mkdir -p domains
COPY ./domains/ ./domains
