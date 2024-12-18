GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
COMMIT=$(shell git rev-parse --short HEAD)

SERVICE_NAME ?= $(shell basename $(CURDIR))

# gRPC Tools
PROTOC := protoc
PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc
PROTO_DIR := ./api/whrmi/v1
PROTO_FILE := $(PROTO_DIR)/*.proto
PROTO_OUT := ./api/whrmi/v1

.PHONY: init build lint install install-grpc generate-proto api

ifeq ($(GOHOSTOS), windows)
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif


init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

# Generate gRPC code from proto files
api:
	$(PROTOC) --proto_path=$(PROTO_DIR) \
			--proto_path=./third_party \
			--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
			--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=$(PROTO_OUT) --grpc-gateway_opt=paths=source_relative \
			$(PROTO_FILE)

# Generate wire_gen
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./internal/di

all: api generate

# api:
# 	protoc --proto_path=./api \
# 	       --proto_path=./third_party \
#  	       --go_out=paths=source_relative:./api \
#  	       --go-http_out=paths=source_relative:./api \
#  	       --go-grpc_out=paths=source_relative:./api \
#  	       --go-errors_out=paths=source_relative:./api \
# 	       --openapi_out=fq_schema_naming=true,default_response=false:. \
# 	       --validate_out=paths=source_relative,lang=go:./api \
# 	       $(API_PROTO_FILES)

# build
build:
# go build -ldflags  '-X github.com/s-yakubovskiy/whereami/cmd/cmd.Version=2.0.1 -X github.com/s-yakubovskiy/whereami/cmd/cmd.Commit=adndf32nd' -o ./bin/ ./... 
	mkdir -p bin/ && go build -ldflags "-X github.com/s-yakubovskiy/whereami/internal/data/zosh.Version=$(VERSION) -X github.com/s-yakubovskiy/whereami/internal/data/zosh.Commit=$(COMMIT)" -o ./bin/ ./...


# lint
lint:
	golangci-lint run --timeout 10m

# install
install: build
	@sudo systemctl stop whereami.service
	@cp ./bin/whereami ~/.local/bin/whereami
	@test -e ~/.config/whereami/config.yaml || cp ./config/config.yaml ~/.config/whereami/
	@sudo systemctl daemon-reload
	@sudo systemctl start whereami.service

run: build
	@export $(shell sed 's/=.*//' ./.env) && \
	source ./.env && \
	./bin/whereami  $(filter-out $@,$(MAKECMDGOALS))

%:
	@:

help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
