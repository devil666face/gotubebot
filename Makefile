PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
GOOS=linux
GOARCH=amd64
CGO_ENABLED=1
LDFLAGS="-w -s"
APP=gotubebot

.PHONY: build
build: ## Build project
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=$(LDFLAGS) -o $(APP) cmd/main/main.go

.PHONY: run
run: ## Run project
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=$(LDFLAGS) -o tmp/$(APP) cmd/main/main.go
	tmp/$(APP)

.PHONY: air
air: ## Run dev server
	go install github.com/cosmtrek/air@latest
	$(GOPATH)/bin/air

.PHONY: .install-linter
.install-linter: # Install linter
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.54.2

.PHONY: lint
lint: .install-linter # Run linter
	$(GOLANGCI_LINT) run ./... --deadline=30m --enable=misspell --enable=gosec --enable=gofmt --enable=goimports --enable=revive 


.PHONY: help
help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
