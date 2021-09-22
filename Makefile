.PHONY: precommit
precommit: format test

.PHONY: test
test:
	go test ./...

.PHONY: test
format:
	go fmt ./...
