include ../../internal/setup/ServiceCommon.mk

RELATIVE_PATH_TO_ROOT = ../..
SERVICE_CONTAINER_NAME = human-service

SERVICE_EXTERNAL_CONNECTION_NETWORK=human-ext-network

create-external-connection-network:
	sudo docker network create --driver bridge ${SERVICE_EXTERNAL_CONNECTION_NETWORK} || true

compose-with-mongo: build-docker create-external-connection-network
	sudo docker compose up

clean:
	sudo docker compose stop

.PHONY: create-external-connection-network