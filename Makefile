lint:
	@echo "Running golangci-lint..."
	@golangci-lint run --timeout 3m --config=.golangci.yml
