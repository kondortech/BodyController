FROM body-controller-monorepo

WORKDIR /domains/nutrition/services/base/ingredient/

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
