include ../../../../internal/setup/CommonVariables.mk
include ../../../../internal/setup/ServiceCommon.mk
include ../../../../internal/setup/AggregationService.mk

DOMAIN_NAME = users

run-docker: build-docker
	sudo docker run -p 8090:8090 --network=${ONLY_BASE_SERVICE_DEBUG_NETWORK} --env-file=./configs/dev/env.config --name=$(SERVICE_CONTAINER_NAME) $(SERVICE_CONTAINER_NAME)