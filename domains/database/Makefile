include ../../internal/setup/Common.mk

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
	migrate -path ./dbmigrations/ -database mongodb://0.0.0.0:27017/BodyController -verbose up

drop-all-migrations-to-local:
	migrate -path ./dbmigrations/ -database mongodb://0.0.0.0:27017/BodyController -verbose down

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