-include .env

SHELL=/bin/bash -o pipefail
GIT_TAG := $(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
PROJECT_NAME := $(shell basename "$(PWD)")
BUILD_VARS_IMPORT_PATH := commerceiq.ai/ticketing/cmd

# Go related variables.
GOBASE := $(shell pwd)
GOFILES := $(wildcard *.go)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags '-X=${BUILD_VARS_IMPORT_PATH}.GitTag=$(GIT_TAG) -X=${BUILD_VARS_IMPORT_PATH}.GitCommit=$(GIT_COMMIT) -extldflags "-static"'

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECT_NAME)-stderr.txt

all: help

## install: Install (for current OS) the compiled binary /usr/local/bin. Run make compile before running this
install: export-path
	@-touch $(STDERR)
	@-rm $(STDERR)
	@chmod +x ticketing
	@mv ticketing /usr/local/bin/
	@printf "\nSuccessfully Installed Ticketing!!\n"

export-path:
	@export PATH=${PATH}:/usr/local/bin/

## compile: download go module dependencies and compile ticketing binary
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-clean 2> $(STDERR)
	@-$(MAKE) -s go-get 2> $(STDERR)
	@-$(MAKE) -s go-build 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-rm $(GOBIN)/$(PROJECT_NAME) 2> /dev/null
	@-$(MAKE) go-clean

## docker: build docker image
docker: docker-build

docker-build-latest:
	docker build . -t ticketing:latest

## docker-build: build docker image
docker-build:
	docker build . -t ticketing:$(GIT_TAG)
## test: run tests
test: go-test

## run go-fmt
fmt: go-fmt

## test-report: generate cov report
test-report:
	@cat test-reports/report.txt | go-junit-report > ./test-reports/report.xml

go-build:
	@echo "  >  Building binary..."
	go build $(LDFLAGS) -o ./$(PROJECT_NAME) $(GO_FILES)

go-build-linux:
	@echo "  >  Building binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build $(LDFLAGS) -o ./$(PROJECT_NAME)-linux $(GO_FILES)

go-generate:
	@echo "  >  Generating dependency files..."
	go generate $(generate)

go-get:
	@echo "  >  Checking if there are any missing dependencies..."
	go get $(get)

go-install:
	go install $(GO_FILES)

go-clean:
	@echo "  >  Cleaning build cache"
	go clean

go-download:
	@echo "  >  Downloading packages"
	go mod download

go-tidy:
	@echo "  >  Cleaning redundant packages"
	go mod tidy

go-fmt:
	@echo "  >  Running go fmt"
	go fmt ./...

go-vet:
	@echo "  >  Running go vet"
	go vet

go-test:
	@echo "  >  Running Tests"
	@mkdir -p test-reports
	go test -v ./... | tee test-reports/report.txt

## help: help for all commands
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

pre-commit:
	@git config --local core.hooksPath hooks

