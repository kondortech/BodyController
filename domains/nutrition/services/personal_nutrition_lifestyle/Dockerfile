FROM body-controller-monorepo

WORKDIR /domains/nutrition/services/base/personal_nutrition_lifestyle/

RUN go build -mod=mod -o service ./*.go

CMD [ "./service" ]
