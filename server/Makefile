.PHONY: run start build test dev

run: ## Run the server (development)
	go run main.go

start: ## Alias for run
	$(MAKE) run

dev: ## Run with live reloading (requires air)
	cd .. && $(shell go env GOPATH)/bin/air

build: ## Build binary
	go build -o ../bin/server main.go

test: ## Run tests (none yet)
	@echo "No tests yet"

