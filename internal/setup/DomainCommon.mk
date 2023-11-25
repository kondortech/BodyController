RELATIVE_PATH_TO_ROOT=../../
SERVICE_DOCKERFILE_PATHS = $(shell find ./ -name 'Dockerfile')

build-service-images: pack-monorepo-in-docker-from-service
	for service_dockerfile_path in $(SERVICE_DOCKERFILE_PATHS); do \
		service_path=$$(dirname $$service_dockerfile_path); \
		cd $$service_path && make build-docker && cd -; \
	done

compose-domain: build-service-images
	sudo docker compose up -d

.PHONY: build-service-images compose-domain