FROM body-controller-base

WORKDIR /services/nutrition/ingredient/

RUN go build -mod=mod ./main/service.go

CMD [ "./service" ]
