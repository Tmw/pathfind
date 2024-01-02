GOTESTSUM_PATH ?= $(shell which gotestsum)
GO_TEST_FLAGS  ?= -v -race -count=1

run:
	@go run cmd/main.go -filename examples/small.txt

test:
	$(if $(GOTESTSUM_PATH), gotestsum --, go test) $(GO_TEST_FLAGS) ./...

lint:
	@golangci-lint run -c .golangci.yml
