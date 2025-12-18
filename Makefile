.PHONY: run start build test

run: ## Run the server (development)
	go run main.go

start: ## Alias for run
	$(MAKE) run

build: ## Build binary
	go build -o bin/server main.go

test: ## Run tests (none yet)
	@echo "No tests yet"

