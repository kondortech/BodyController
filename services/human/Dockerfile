FROM body-controller-monorepo

WORKDIR /services/human/

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
