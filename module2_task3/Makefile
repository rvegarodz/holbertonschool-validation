.PHONY: build run stop clean test unit-tests integration-tests

## Variables
BINARY_NAME := awesome-api
LOG_FILE := $(BINARY_NAME).log
PID_FILE := $(BINARY_NAME).pid
UNIT_FILE := coverage-units.out
COVERAGE_FILE := coverage-integrations.out

build: ## Compile the source code of the application to a binary.
	golangci-lint run
	go get github.com/gorilla/mux
	make integration-tests
	go build -o $(BINARY_NAME)

run: ## Run the application by executing the binary file.
	./$(BINARY_NAME) > $(LOG_FILE) 2>&1 & \
	echo $$! > $(PID_FILE)

stop: ## Stop the application and remove PID file.
	-kill `cat $(PID_FILE)` || true
	-rm $(PID_FILE) || true

clean: ## Stop the application. Delete the binary and log file.
	make stop
	-rm $(BINARY_NAME) || true
	-rm $(LOG_FILE) || true
	-rm $(UNIT_FILE) || true
	-rm $(COVERAGE_FILE) || true

test: ## Test the server.
	make unit-tests
	make integration-tests

unit-tests: ## Run the Unit Tests.
	go test -short -coverprofile=$(UNIT_FILE) -run Test_HelloHandler

integration-tests: ## Run the Integration Tests.
	go test -short -coverprofile=$(COVERAGE_FILE) -run Test_server

lint: ## Check the Go source file.
	golangci-lint run

help: ## Prints out the list of targets and their usage.
	@awk 'BEGIN {FS = ":.*?# "} /^[a-zA-Z_-]+:.*?# / {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
