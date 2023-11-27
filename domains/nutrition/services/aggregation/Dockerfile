FROM body-controller-monorepo

WORKDIR /domains/users/services/aggregation/

RUN go build -mod=mod -o service ./main/*

CMD [ "./service" ]
