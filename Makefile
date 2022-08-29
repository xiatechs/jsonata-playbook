.PHONY: help

CONTAINERS = $(shell docker ps -a -q)
VOLUMES = $(shell docker volume ls -q)

help: ## Show this help menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## test the go logic
	@go test -race -cover ./...

start: ## start the docker container
	@docker build -t jsonatagui . && docker run -dp 8050:8050 jsonatagui

unbuild: ## stop & remove docker containers
	docker rm -f $(CONTAINERS)
	docker volume rm $(VOLUMES)