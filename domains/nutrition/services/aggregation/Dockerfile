FROM body-controller-monorepo

WORKDIR /domains/nutrition/services/aggregation/

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
