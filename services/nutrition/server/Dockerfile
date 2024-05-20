FROM body-controller-monorepo

WORKDIR /services/nutrition/server

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
