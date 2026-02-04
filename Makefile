GO ?= go
APP_NAME ?= tetris
CMD_DIR ?= ./cmd
BIN_DIR ?= ./build
BIN_PATH := $(BIN_DIR)/$(APP_NAME)

.DEFAULT_GOAL := help

.PHONY: help run build test test-race fmt fmt-check vet tidy clean check

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "%-12s %s\n", $$1, $$2}'

run: ## Run the application
	$(GO) run $(CMD_DIR)

build: ## Build binary into ./build
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_PATH) $(CMD_DIR)

test: ## Run unit tests
	$(GO) test ./...

test-race: ## Run tests with race detector
	$(GO) test -race ./...

fmt: ## Format all Go files
	$(GO) fmt ./...

fmt-check: ## Fail if formatting is needed
	@test -z "$$(gofmt -l .)" || (echo "Run 'make fmt' to format files" && gofmt -l . && exit 1)

vet: ## Run go vet
	$(GO) vet ./...

tidy: ## Tidy go.mod and go.sum
	$(GO) mod tidy

check: fmt-check vet test ## Run basic CI checks

clean: ## Remove build artifacts
	rm -rf $(BIN_DIR)
