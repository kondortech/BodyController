PROTO_FILES = $(shell find ./ -name '*.proto')

generate-all-protos:
	for file in $(PROTO_FILES); do \
		proto_dir=$$(dirname "$$file"); \
		proto_basename=$$(basename "$$file"); \
		echo "generating proto for $$proto_dir $$proto_basename"; \
		cd $$proto_dir; \
		protoc --proto_path=. --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $$proto_basename; \
		cd -; \
		echo "generated proto for $$proto_dir $$proto_basename"; \
	done

pack-monorepo-in-docker: generate-all-protos
	sudo docker build . -t body-controller-monorepo

.PHONY: pack-monorepo-in-docker generate-all-protos
