.PHONY: env
env:  ## 环境配置
	@echo "protoc version is: "
	@protoc --version
	@echo "install dependence package"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: build
build:  ## 构建
	@echo "build start"
	@protoc -I . \
         -I ../third_party \
         --go_out=paths=source_relative:.  \
         --go-grpc_out==paths=source_relative:.  \
         --grpc-gateway_out=paths=source_relative:.  \
         --grpc-gateway_opt logtostderr=true \
         --grpc-gateway_opt generate_unbound_methods=true \
          --openapi_out=fq_schema_naming=true,default_response=false,paths=source_relative:. \
         ./api/hello.proto
	@if (( $$? == 0 )) ;then echo "build successful"; else echo "build failed"; fi

.PHONY: cer
cer:  ## 生成证书
	@echo "生成证书"
	@openssl version
	@openssl genrsa -out server.key 4096
	@openssl req -new -key server.key -out server.csr -subj "/C=US/ST=New York/L=New York/O=Example/CN=example.com"
	@openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt


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