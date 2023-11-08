.PHONY: build
build:  ## 生成api
	@echo "generate api of http server and rpc server and api yaml file"
	@protoc --proto_path=. \
     	    --proto_path=../../../third_party \
      	    --go_out=paths=source_relative:. \
      	    --go-http_out=paths=source_relative:. \
      	    --go-grpc_out=paths=source_relative:. \
     	    --openapi_out=fq_schema_naming=true,default_response=false:. \
     	    ./hello.proto

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