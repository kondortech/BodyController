create-ts-output-dir:
	mkdir -p ./src/generated/services/nutrition


generate-client-definitions: create-ts-output-dir
	npx swagger-typescript-api \
		--union-enums \
		-p ../generated/openapiv2/services/nutrition/proto/api.swagger.json \
		-o ./src/generated/services/nutrition \
		-n api.ts; \

.PHONY: create-ts-output-dir generate-client-definitions