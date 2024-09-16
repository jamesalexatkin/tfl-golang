lint:
	@echo Running linter
	golangci-lint cache clean
	golangci-lint run --config=./.golangci.yml

test:
	@echo Running unit tests
	go clean -testcache && \
	go test -short ./...

test-integration:
	@echo Running integration tests
	go clean -testcache && \
	go test -run Integration ./...

test-all:
	@echo Running unit and integration tests
	go clean -testcache && \
	go test ./...