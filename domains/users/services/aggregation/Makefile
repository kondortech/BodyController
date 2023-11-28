include ../../../../internal/setup/CommonVariables.mk
include ../../../../internal/setup/ServiceCommon.mk
include ../../../../internal/setup/AggregationService.mk

DOMAIN_NAME = users
SERVICE_PORT = 10000

BASE_AUTH_SERVICE_NAME = users-base-auth
BASE_AUTH_SERVICE_PORT = 10001

run-docker: build-docker
	sudo docker run \
		--expose ${SERVICE_PORT} \
		-p ${SERVICE_PORT}:${SERVICE_PORT} \
		--network=${ONLY_BASE_SERVICE_DEBUG_NETWORK} \
		-e SERVICE_PORT=${SERVICE_PORT} \
		-e BASE_AUTH_IP=${BASE_AUTH_SERVICE_NAME} \
		-e BASE_AUTH_PORT=${BASE_AUTH_SERVICE_PORT} \
		--name=$(SERVICE_CONTAINER_NAME) $(SERVICE_CONTAINER_NAME)