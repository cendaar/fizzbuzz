.PHONY: build
build:
	go build -o fizzbuzz cmd/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run: build
	./fizzbuzz

.PHONY: format
format:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run