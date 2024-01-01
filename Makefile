run:
	@go run cmd/main.go -filename examples/small.txt

test:
	@gotestsum
	# @go test ./...
