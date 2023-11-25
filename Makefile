PROTO_FILES = $(shell find ./ -name '*.proto')
PROJECT_ROOT = $(shell pwd)

generate-all-protos:
	for file in $(PROTO_FILES); do \
		filepath_relative_to_project_root=$$(echo $$file | sed 's/^..//'); \
		protoc --proto_path=${PROJECT_ROOT} --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $$filepath_relative_to_project_root; \
	done

pack-monorepo-in-docker: generate-all-protos
	sudo docker build . -t body-controller-monorepo

.PHONY: pack-monorepo-in-docker generate-all-protos cur
