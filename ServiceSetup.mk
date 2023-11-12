# This file should be included ONLY in makefiles of services
SERVICE_TO_ROOT_RELATIVE = ../../..
PROJECT_ROOT = $(shell cd $(SERVICE_TO_ROOT_RELATIVE) && pwd && cd ~-)

.PHONY: prepare-monorepo-from-service build-containerized start-containerized
prepare-monorepo-from-service:
	cd $(PROJECT_ROOT) && make prepare-monorepo && cd -

build-containerized: prepare-monorepo-from-service
	sudo docker build . -t $(DOMAIN_NAME)/$(SERVICE_NAME)

start-containerized: build-containerized
	sudo docker run $(DOMAIN_NAME)/$(SERVICE_NAME)
