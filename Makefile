GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

#SERVICE_NAME ?= $(shell basename $(CURDIR)) # this can be changed if service name doesn't match with current directory
SERVICE_NAME ?=whereami# this can be changed if service name doesn't match with current directory

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env

.PHONY: build
# build
build:
	@mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...


.PHONY: lint
# lint
lint:
	golangci-lint run --timeout 10m

.PHONY: install
# install
install: build
	@sudo systemctl stop whereami.service
	@cp ./bin/whereami ~/.local/bin/whereami
	@cp ./config/config.yaml ~/.config/whereami/
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
