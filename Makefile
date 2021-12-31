
APP_NAME := "app"

.PHONY: build
build:
	go build -o $(APP_NAME) main.go

.PHONY: proto
proto:
	protoc -I pb \
	--go_out api \
	--go_opt paths=source_relative \
	--go-grpc_out api \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out api \
	--grpc-gateway_opt paths=source_relative \
	pb/**/**/*.proto