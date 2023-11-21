MONGODB_CONTAINER_NAME = body-controller-mongo-db
SERVICE_TO_ROOT_RELATIVE = ../../../../..
PROJECT_ROOT = $(shell cd $(SERVICE_TO_ROOT_RELATIVE) && pwd && cd -)

# TODO construct DOMAIN_NAME and SERVICE_NAME with assumptions that this file is in the base service root dir
# Every Makefile for now will have custom variables DOMAIN_NAME and SERVICE_NAME
BASE_SERVICE_DOCKER_CONTAINER_NAME = $(DOMAIN_NAME)/base-$(SERVICE_NAME)

# TODO Generalize for any Makefile
generate-all-protos:
	cd $(PROJECT_ROOT) && make generate-all-protos && cd -

# TODO Generalize for any Makefile
pack-monorepo-in-docker-from-service: generate-all-protos
	cd $(PROJECT_ROOT) && make pack-monorepo-in-docker && cd -

build-docker: pack-monorepo-in-docker-from-service
	sudo docker build --no-cache --progress=plain -t $(BASE_SERVICE_DOCKER_CONTAINER_NAME) .

run-docker: build-docker
	sudo docker network create ${DEFAULT_DEBUG_NETWORK}
	sudo docker network connect ${DEFAULT_DEBUG_NETWORK} ${MONGODB_CONTAINER_NAME}
	sudo docker run -p ${DEFAULT_BASE_SERVICE_PORT}:8080 --network=${DEFAULT_DEBUG_NETWORK} $(BASE_SERVICE_DOCKER_CONTAINER_NAME)

clean-base-service-docker:
	sudo docker stop ${BASE_SERVICE_DOCKER_CONTAINER_NAME}
	sudo docker network disconnect ${DEFAULT_DEBUG_NETWORK} ${BASE_SERVICE_DOCKER_CONTAINER_NAME}
	sudo docker rm ${BASE_SERVICE_DOCKER_CONTAINER_NAME}

clean-debug-network:
	sudo docker network disconnect ${DEFAULT_DEBUG_NETWORK} ${MONGODB_CONTAINER_NAME}
	sudo docker network remove ${DEFAULT_DEBUG_NETWORK}

clean-docker: clean-base-service-docker clean-debug-network

run-local: generate-all-protos
	go run -mod=mod ./main/*

clean-local:
	rm service

.PHONY: generate-all-protos pack-monorepo-in-docker-from-service build-docker start-docker clean-docker build-local run-local clean-local

