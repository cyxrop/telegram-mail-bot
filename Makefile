include ./docker/.env
DB_CONN="host=localhost port=5432 user=postgres password=postgres dbname=homework_2 sslmode=disable"

.PHONY: all
all: run

.PHONY: proto-gen
proto-gen:
	protoc -I ./proto \
	-I ${GOPATH}/src/github.com/googleapis \
	--go_out=pkg --go_opt=paths=source_relative \
	--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out ./pkg \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--swagger_out=allow_merge=true,merge_file_name=api:docs \
	proto/api/users.proto

.PHONY: migrate
migrate:
	goose -dir=./migrations postgres "user=postgres password=postgres dbname=homework_2 sslmode=disable" up

.PHONY: run
run:
	CRYPT_KEY=${CRYPT_KEY} \
	TG_TOKEN=${TG_TOKEN} \
	DB_CONN=${DB_CONN} \
	go run ./cmd/app/main.go