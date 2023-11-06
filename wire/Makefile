.PHONY: init
init:  ## 工具安装
	@echo "install gozz"
	@echo "参考文档地址: https://go-zing.github.io/gozz/guide/getting-started.html#annotation-syntax"
	@go install github.com/go-zing/gozz@latest

.PHONY: wire
wire:  ## 依赖注入
	@gozz  run -p "wire" ./...

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