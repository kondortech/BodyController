include ../../../../internal/setup/CommonVariables.mk
include ../../../../internal/setup/ServiceCommon.mk
include ../../../../internal/setup/AggregationService.mk

DOMAIN_NAME = nutrition
SERVICE_PORT = 20000

BASE_INGREDIENT_SERVICE_NAME = nutrition-base-ingredient
BASE_INGREDIENT_SERVICE_PORT = 20001

run-docker: build-docker
	sudo docker run \
		--expose ${SERVICE_PORT} \
		-p ${SERVICE_PORT}:${SERVICE_PORT} \
		--network=${ONLY_BASE_SERVICE_DEBUG_NETWORK} \
		-e SERVICE_PORT=${SERVICE_PORT} \
		-e BASE_INGREDIENT_IP=${BASE_INGREDIENT_SERVICE_NAME} \
		-e BASE_INGREDIENT_PORT=${BASE_INGREDIENT_SERVICE_PORT} \
		--name=$(SERVICE_CONTAINER_NAME) $(SERVICE_CONTAINER_NAME)