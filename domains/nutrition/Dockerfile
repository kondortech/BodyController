FROM body-controller-monorepo

WORKDIR /domains/nutrition/

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
