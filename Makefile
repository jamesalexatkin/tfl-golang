lint:
	@echo Running linter
	golangci-lint cache clean
	golangci-lint run --config=./.golangci.yml

test:
	@echo Running tests
	go clean -testcache && \
	go test ./...