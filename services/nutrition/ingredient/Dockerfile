FROM golang:1.21
SHELL ["/bin/bash", "-c"]

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download


RUN mkdir -p main
COPY src/* ./src/

RUN mkdir -p proto
COPY proto/* ./proto/

RUN go build -mod=mod ./src/service.go

CMD [ "./service" ]
