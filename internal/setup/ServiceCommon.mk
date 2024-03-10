PROJECT_ROOT = $(shell cd $(RELATIVE_PATH_TO_ROOT) && pwd && cd -)
MONGODB_CONTAINER_NAME = body-controller-mongodb-container

generate-all-protos:
	cd $(PROJECT_ROOT) && make generate-all-protos && cd -

pack-monorepo-in-docker-from-service: generate-all-protos
	cd $(PROJECT_ROOT) && make pack-monorepo-in-docker && cd -

build-docker: pack-monorepo-in-docker-from-service
	sudo docker build --progress=plain -t $(SERVICE_CONTAINER_NAME) .

run-local: generate-all-protos
	go run -mod=mod ./cmd.go

clean-local:
	rm service

.PHONY: generate-all-protos pack-monorepo-in-docker-from-service build-docker run-local clean-local