SHELL := /bin/sh
export PATH := $(PATH)

all: fmt lintfix tidy test clean build

debug:
	@echo "PATH: $(PATH)"
	@echo "GOPATH: $(shell go env GOPATH)"
	@which go
	@which gofumpt
	@which golangci-lint

clean:
	rm -rf ./dist

# Test
test:
	go test ./...

tidy:
	go mod tidy

# Format code
fmt:
	gofumpt -l -w .

# Run linter
lint: fmt
	golangci-lint run

# Run linter and fix
lintfix: fmt
	golangci-lint run --fix

build: clean
	mkdir ./dist
	go build -o ./dist/gmi2html ./cmd/gmi2html

build-gccgo: clean
	go build -compiler=gccgo -o ./dist/gmi2html ./cmd/gmi2html

show-updates:
	go list -m -u all

update:
	go get -u all

update-patch:
	go get -u=patch all
