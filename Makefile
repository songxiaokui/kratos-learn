GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
init: ## 项目依赖初始化
	@echo "download proto dependence such as proto-gen-go and kratos etc."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@go install github.com/google/wire/cmd/wire@latest

.PHONY: generate
generate:  ## 依赖链自动生成
	@echo "init google wire and dependency injection"
	@go mod tidy
	@go get github.com/google/wire/cmd/wire@latest
	@go generate ./...


.PHONY: config
config:  ## 配置管理
	@echo "config manager code generate"

.PHONY: api
api:  ## api生成
	@echo "generate api of http and rpc"

.PHONY: all
all: ## 构建所有
	@echo "generate all server"

.PHONY: build
build:  ## 打包应用
	@echo "build server"

.PHONY: run
run:  ## 运行应用
	@echo "运行应用"

help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help