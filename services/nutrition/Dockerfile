FROM body-controller-monorepo

WORKDIR /services/nutrition/

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
