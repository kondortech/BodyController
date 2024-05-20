FROM body-controller-monorepo

WORKDIR /services/gateway/server

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
