ifndef VERBOSE
	MAKEFLAGS += --silent
endif

TARGET = api-server
GOTARGET = github.com/peaklyio/$(TARGET)
REGISTRY ?= peaklyio
IMAGE = $(REGISTRY)/$(TARGET)
DIR := ${CURDIR}
DOCKER ?= docker
PKGS=$(shell go list ./... | grep -v /vendor)
GIT_VERSION ?= $(shell git describe --always --dirty)
IMAGE_VERSION ?= $(shell git describe --always --dirty)
IMAGE_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD | sed 's/\///g')
GIT_REF = $(shell git rev-parse --short=8 --verify HEAD)

FMT_PKGS=$(shell go list -f {{.Dir}} ./... | grep -v vendor | tail -n +2)

default: compile

all: compile install

push: ## Push to the docker registry
	$(DOCKER) push $(REGISTRY)/$(TARGET):$(GIT_REF)
	$(DOCKER) push $(REGISTRY)/$(TARGET):latest

clean: ## Clean the docker images
	rm -f $(TARGET)
	$(DOCKER) rmi $(REGISTRY)/$(TARGET) || true

container: ## Build the docker container
	$(DOCKER) build \
		-t $(REGISTRY)/$(TARGET):$(IMAGE_VERSION) \
		-t $(REGISTRY)/$(TARGET):$(IMAGE_BRANCH) \
		-t $(REGISTRY)/$(TARGET):$(GIT_REF) \
	    -t $(REGISTRY)/$(TARGET):latest \
		.

run: ## Run the api-server in a container
	$(DOCKER) run $(REGISTRY)/$(TARGET):$(IMAGE_VERSION)


compile: ## Compile the binary into bin/api-server
	go build -o bin/api-server main.go

install: ## Create the api-server executable in $GOPATH/bin directory.
	install -m 0755 bin/api-server ${GOPATH}/bin/api-server

gofmt: install-tools ## Go fmt your code
	echo "Fixing format of go files..."; \
	for package in $(FMT_PKGS); \
	do \
		gofmt -w $$package ; \
		goimports -l -w $$package ; \
	done

check-headers: ## Check if the headers are valid. This is ran in CI.
	./scripts/check-header.sh

.PHONY: test
test: ## Run the INTEGRATION TESTS. This will create cloud resources and potentially cost money.
	go test -timeout 20m -v $(PKGS)

.PHONY: check-code
check-code: install-tools ## Run code checks
	PKGS="${FMT_PKGS}" GOFMT="gofmt" GOLINT="golint" ./scripts/ci-checks.sh

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install-tools
install-tools:
	GOIMPORTS_CMD=$(shell command -v goimports 2> /dev/null)
ifndef GOIMPORTS_CMD
	go get golang.org/x/tools/cmd/goimports
endif
	GOLINT_CMD=$(shell command -v golint 2> /dev/null)
ifndef GOLINT_CMD
	go get github.com/golang/lint/golint
endif