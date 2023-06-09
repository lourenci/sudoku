.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Test the project
	go test ./...

.PHONY: coverage
coverage: ## Generate the code coverage report
	go test -coverprofile=coverage.txt -covermode=atomic ./...

open-coverage: coverage ## Generate the code coverage report and open in the browser
	go tool cover -html=coverage.txt
