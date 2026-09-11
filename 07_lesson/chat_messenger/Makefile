.PHONY: openapi proto generate clean lint

build:

openapi:
	oapi-codegen \
	-config api/openapi/oapi-codegen.yaml \
	api/openapi/messenger.swagger.yaml

proto:
	mkdir -p gen
	protoc \
		--proto_path=api/proto \
		--go_out=gen \
        --go_opt=paths=source_relative \
        --go-grpc_out=gen \
        --go-grpc_opt=paths=source_relative \
        api/proto/messenger/v1/messenger.proto

generate: proto openapi

clean:
	rm -rf gen/

lint:
	golangci-lint run