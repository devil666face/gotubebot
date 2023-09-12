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
	~/go/bin/air

.PHONY: lint
lint: ## Run linter
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ~/go/bin v1.54.2
	~/go/bin/golangci-lint run --deadline=30m --enable=misspell --enable=gosec --enable=gofmt --enable=goimports --enable=revive 
	~/go/bin/golangci-lint run

.PHONY: help
help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
