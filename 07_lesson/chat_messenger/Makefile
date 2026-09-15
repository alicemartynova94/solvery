DB_USER=postgres
DB_PASSWORD?=pass_example
DB_HOST=localhost
DB_PORT=5432
DB_NAME=chat_messenger

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS=./internal/database/migrations

.PHONY: build run openapi proto generate clean lint migrate-create migrate-up migrate-down

build:
	go build -o bin/messenger ./cmd/messenger

run:
	go run ./cmd/messenger

openapi:
	oapi-codegen \
	-config api/openapi/oapi-codegen.yaml \
	api/openapi/messenger.swagger.yaml

proto:
	mkdir -p pkg
	protoc \
		--proto_path=api/proto \
		--go_out=pkg \
        --go_opt=paths=source_relative \
        --go-grpc_out=pkg \
        --go-grpc_opt=paths=source_relative \
        api/proto/messenger/v1/messenger.proto

generate: proto openapi

clean:
	rm -rf pkg/

lint:
	golangci-lint run

migrate-create:
	goose -dir $(MIGRATIONS) postgres "$(DB_URL)" create $(name) sql

migrate-up:
	goose -dir $(MIGRATIONS) postgres "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS) postgres "$(DB_URL)" down