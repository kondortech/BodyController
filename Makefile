.PHONY: prepare-monorepo
prepare-monorepo:
	sudo docker build . -t body-controller-monorepo

generate-models:
	sudo echo $(shell find ./ -name '*.proto')
	sudo protoc \
        --go_out protos/gen \
        --go_opt paths=source_relative \
        --go-grpc_out protos/gen \
        --go-grpc_opt paths=source_relative \
        $(shell find ./ -name '*.proto')

PROTO_FILES = $(shell find ./ -name '*.proto')

names:
	parentdir=./
	for name in $(shell find ./ -name '*.proto'); do \
		dirname $$name; \
	done