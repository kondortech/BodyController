FROM body-controller-monorepo

WORKDIR /domains/nutrition/services/base/nutrition_lifestyle_template/

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
