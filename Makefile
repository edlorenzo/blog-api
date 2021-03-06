#for newmand command HOST must be non-localhost 

HOST=localhost
#HOST=192.168.56.103
PORT=5007
CONTAINER_PORT=5007
#export APIURL=http://192.168.56.103:5007/api
export APIURL=http://$(HOST):$(PORT)/api

GOOS=linux
GOARCH=amd64
APP=fiber-rw
APP_STATIC=$(APP)-static
LDFLAGS="-w -s -extldflags=-static"

USERNAME=u$(shell date +%s)
EMAIL=$(USERNAME)@mail.com
PASSWORD=password
NEWMAN_URL=Blog-API.postman_collection.json

tidy:
	go mod tidy -compat=1.17

help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

download: ## Download go dependency 
	go mod download

clear-db: ## Remove old database
	rm -f ./database/blog.db

swagger-init: ## Swagger Init.
	swag init

generate: ## Generate swagger docs. Required https://github.com/arsmn/fiber-swagger 
	go generate .	

build: ## Build project with dynamic library(see shared lib with "ldd <your_file>") 
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -race -o $(APP) . 


test: ## Run unit test without race detection 
	go test -v ./...

test-race: ## Run unit test without race detection
	go test -v -race  ./...

