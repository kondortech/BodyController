RELATIVE_PATH_TO_ROOT = ../../../../..

# TODO construct DOMAIN_NAME and SERVICE_NAME with assumptions that this file is in the base service root dir
# Every Makefile for now will have custom variables DOMAIN_NAME and SERVICE_NAME
SERVICE_CONTAINER_NAME = $(DOMAIN_NAME)-base-$(SERVICE_NAME)

MONGODB_IP=body-controller-mongodb-container
MONGODB_PORT=27017

run-docker: build-docker
	sudo docker run \
		-p ${SERVICE_PORT}:${SERVICE_PORT} \
		--expose ${SERVICE_PORT} \
		--network=${ONLY_BASE_SERVICE_DEBUG_NETWORK} \
		-e SERVICE_PORT=${SERVICE_PORT} \
		-e MONGODB_IP=${MONGODB_IP} \
		-e MONGODB_PORT=${MONGODB_PORT} \
		--name=$(SERVICE_CONTAINER_NAME) $(SERVICE_CONTAINER_NAME)

clean-docker:
	sudo docker container prune

.PHONY: run-docker clean-docker