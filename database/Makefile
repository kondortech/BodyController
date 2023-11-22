include ../../internal/setup/Common.mk

MIGRATIONS_DIRECTORY = ./dbmigrations/

MONGODB_CONTAINER_NAME = body-controller-mongo-db
MONGODB_PORT = 27017
MONGODB_VOLUME_NAME = mongodbdata

MONGODB_OBSERVER_CONTAINER_NAME = body-controller-mongo-db-observer
MONGODB_OBSERVER_PORT = 27000
MONGODB_OBSERVER_NETWORK_NAME = mongodb-observer-network

create-volume:
	sudo docker volume create ${MONGODB_VOLUME_NAME}

start-mongodb:
	sudo docker run -d \
		-p ${MONGODB_PORT}:27017 \
		-v ${MONGODB_VOLUME_NAME}:/data/db \
		--name ${MONGODB_CONTAINER_NAME} \
		mongo:latest

start-mongodb-observer:
	sudo docker network create ${MONGODB_OBSERVER_NETWORK_NAME}
	sudo docker network connect ${MONGODB_OBSERVER_NETWORK_NAME} ${MONGODB_CONTAINER_NAME}
	sudo docker run -d \
		-e ME_CONFIG_MONGODB_SERVER=${MONGODB_CONTAINER_NAME} \
		--network ${MONGODB_OBSERVER_NETWORK_NAME} \
		-p ${MONGODB_OBSERVER_PORT}:8081 \
		--name ${MONGODB_OBSERVER_CONTAINER_NAME} \
		mongo-express

start-mongodb-with-observer: start-mongodb start-mongodb-observer

apply-all-migrations-to-local:
	migrate -path ${MIGRATIONS_DIRECTORY} -database mongodb://0.0.0.0:${MONGODB_PORT}/${DATABASE_NAME} -verbose up

drop-all-migrations-to-local:
	migrate -path ${MIGRATIONS_DIRECTORY} -database mongodb://0.0.0.0:${MONGODB_PORT}/${DATABASE_NAME} -verbose down

clean-observer:
	sudo docker network disconnect ${MONGODB_OBSERVER_NETWORK_NAME} ${MONGODB_CONTAINER_NAME}
	sudo docker network disconnect ${MONGODB_OBSERVER_NETWORK_NAME} ${MONGODB_OBSERVER_CONTAINER_NAME}
	sudo docker stop ${MONGODB_OBSERVER_CONTAINER_NAME}
	sudo docker rm ${MONGODB_OBSERVER_CONTAINER_NAME}
	sudo docker network remove ${MONGODB_OBSERVER_NETWORK_NAME}

clean-db:
	sudo docker stop ${MONGODB_CONTAINER_NAME}
	sudo docker rm ${MONGODB_CONTAINER_NAME}

clean: clean-observer clean-db

.PHONY: start-mongodb start-mongodb-observer start-mongodb-with-observer apply-all-migrations-to-local drop-all-migrations-to-local clean-db clean-observer clean

TST_MONGODB_CONTAINER_NAME = tst-body-controller-mongo-db
TST_MONGODB_PORT = 26017

TST_MONGODB_OBSERVER_CONTAINER_NAME = tst-body-controller-mongo-db-observer
TST_MONGODB_OBSERVER_PORT = 26000
TST_MONGODB_OBSERVER_NETWORK_NAME = tst-mongodb-observer-network

start-mongodb-tst:
	sudo docker run -d \
		-p ${TST_MONGODB_PORT}:27017 \
		--name ${TST_MONGODB_CONTAINER_NAME} \
		mongo:latest

start-mongodb-observer-tst:
	sudo docker network create ${TST_MONGODB_OBSERVER_NETWORK_NAME}
	sudo docker network connect ${TST_MONGODB_OBSERVER_NETWORK_NAME} ${TST_MONGODB_CONTAINER_NAME}
	sudo docker run -d \
		-e ME_CONFIG_MONGODB_SERVER=${TST_MONGODB_CONTAINER_NAME} \
		--network ${TST_MONGODB_OBSERVER_NETWORK_NAME} \
		-p ${TST_MONGODB_OBSERVER_PORT}:8081 \
		--name ${TST_MONGODB_OBSERVER_CONTAINER_NAME} \
		mongo-express

start-mongodb-with-observer-tst: start-mongodb-tst start-mongodb-observer-tst

apply-all-migrations-to-local-tst:
	migrate -path ${MIGRATIONS_DIRECTORY} -database mongodb://0.0.0.0:${TST_MONGODB_PORT}/${DATABASE_NAME} -verbose up

drop-all-migrations-to-local-tst:
	migrate -path ${MIGRATIONS_DIRECTORY} -database mongodb://0.0.0.0:${TST_MONGODB_PORT}/${DATABASE_NAME} -verbose down

clean-observer-tst:
	sudo docker network disconnect ${TST_MONGODB_OBSERVER_NETWORK_NAME} ${TST_MONGODB_CONTAINER_NAME}
	sudo docker network disconnect ${TST_MONGODB_OBSERVER_NETWORK_NAME} ${TST_MONGODB_OBSERVER_CONTAINER_NAME}
	sudo docker stop ${TST_MONGODB_OBSERVER_CONTAINER_NAME}
	sudo docker rm ${TST_MONGODB_OBSERVER_CONTAINER_NAME}
	sudo docker network remove ${TST_MONGODB_OBSERVER_NETWORK_NAME}

clean-db-tst:
	sudo docker stop ${TST_MONGODB_CONTAINER_NAME}
	sudo docker rm ${TST_MONGODB_CONTAINER_NAME}

clean-tst: clean-observer-tst clean-db-tst

.PHONY: start-mongodb-tst start-mongodb-observer-tst start-mongodb-with-observer=tst apply-all-migrations-to-local-tst drop-all-migrations-to-local-tst clean-db-tst clean-observer-tst clean-tst
