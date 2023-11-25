MONGODB_PORT=27017
MONGODB_NETWORK=mongodb-network
MONGODB_VOLUME_NAME=mongodb-data-volume
MONGODB_OBSERVER_NETWORK_NAME=mongodb-observer-network

DATABASE_NAME = BodyController
MIGRATIONS_DIRECTORY = ./dbmigrations/

create-mongodb-volume:
	sudo docker volume create ${MONGODB_VOLUME_NAME}

create-mongodb-network:
	sudo docker network create --driver bridge ${MONGODB_NETWORK} || true

run: create-mongodb-volume create-mongodb-network
	sudo docker compose up -d

clean:
	sudo docker compose stop

apply-all-migrations-local:
	migrate -path ${MIGRATIONS_DIRECTORY} -database mongodb://0.0.0.0:${MONGODB_PORT}/${DATABASE_NAME} -verbose up

drop-all-migrations-local:
	migrate -path ${MIGRATIONS_DIRECTORY} -database mongodb://0.0.0.0:${MONGODB_PORT}/${DATABASE_NAME} -verbose down

.PHONY: create-mongodb-volume create-mongodb-network apply-all-migrations-local drop-all-migrations-local