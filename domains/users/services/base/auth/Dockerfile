FROM body-controller-monorepo

WORKDIR /domains/users/services/base/auth/

RUN go build -mod=mod -o service ./main/*

CMD [ "./service" ]
