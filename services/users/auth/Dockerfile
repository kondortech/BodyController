FROM body-controller-monorepo

WORKDIR /services/users/auth/

RUN go build -mod=mod ./main/service.go

CMD [ "./service" ]
