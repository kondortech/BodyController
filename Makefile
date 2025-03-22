PROTO_FILES = $(shell find . -type f -name '*.proto' ! -path "*/node_modules/*")
PROJECT_ROOT = $(shell pwd)
GENERATED_DIR_NAME = generated

generate-output-openapi-directory:
	mkdir -p ./generated/openapiv2

generate-all-protos: generate-output-openapi-directory
	for file in $(PROTO_FILES); do \
		filepath_relative_to_project_root=$$(echo $$file | sed 's/^..//'); \
		protoc $$filepath_relative_to_project_root \
			--proto_path=${PROJECT_ROOT} \
			--go_out=. \
			--go-grpc_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative \
			--grpc-gateway_out . \
			--grpc-gateway_opt paths=source_relative \
			--grpc-gateway_opt generate_unbound_methods=true \
			--openapiv2_out generated/openapiv2; \
	done

pack-monorepo-for-google-artifactory-registry: generate-all-protos
	docker build -t body-controller-monorepo .
	docker tag body-controller-monorepo europe-west10-docker.pkg.dev/weighty-nation-434312-k6/body-controller-monorepo/test-image
	docker push europe-west10-docker.pkg.dev/weighty-nation-434312-k6/body-controller-monorepo/test-image

load-monorepo-docker-image-to-kind: generate-all-protos
	docker build -t body-controller-monorepo .
	kind load docker-image body-controller-monorepo --name body-controller-local-dev



.PHONY: pack-monorepo-for-google-artifactory-registry generate-all-protos generate-output-openapi-directory
