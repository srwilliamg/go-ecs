
PROJECT_NAME := "gopi"

.PHONY: all fmt tidy help lint test run build clean help build-docker-local docker-run-local start-db

all: fmt tidy lint build test

fmt:
	go fmt ./...

tidy: ## Runs go mod tidy to update dependencies
	go mod tidy

lint: ## Runs linter
	./bin/lint.sh;

test: ## Runs unit tests
	./bin/test.sh;

run: ## Runs the application
	@DB_HOST=localhost APP_ENV=development DB_PORT=5432 DB_USERNAME=user DB_PASSWORD=pass DB_NAME=gopi-db go run main.go

build: clean ## Builds the binary file
	mkdir -p build && go build -v main.go > build/${PROJECT_NAME}

clean: ## Removes previous build
	@rm -rf build

build-docker: ## Build the docker image locally
	@docker build -t ${PROJECT_NAME}:image .

run-docker: ## Runs the locally build docker image (forwards port 8080)
	@docker run --rm -p="8080:8080" --name ${PROJECT_NAME}-container ${PROJECT_NAME}:image

stop-docker: 
	@docker stop ${PROJECT_NAME}-container

start-compose: 
	@docker-compose up

stop-compose: 
	@docker-compose down

clean-compose: 
	@docker rm go-db

start-db:
	@docker compose up postgres_db -d

help: ## Displays this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

