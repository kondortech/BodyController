RELATIVE_PATH_TO_ROOT = ../../../../..

# TODO construct DOMAIN_NAME and SERVICE_NAME with assumptions that this file is in the base service root dir
# Every Makefile for now will have custom variables DOMAIN_NAME and SERVICE_NAME
SERVICE_CONTAINER_NAME = $(DOMAIN_NAME)-base-$(SERVICE_NAME)

run-docker: build-docker
	sudo docker run --expose ${DEFAULT_BASE_SERVICE_PORT}:8080 --network=${ONLY_BASE_SERVICE_DEBUG_NETWORK} --env-file=./configs/dev/env.config --name=$(SERVICE_CONTAINER_NAME) $(SERVICE_CONTAINER_NAME)

clean-docker:
	sudo docker container prune

.PHONY: run-docker clean-docker