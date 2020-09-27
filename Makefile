ifeq ($(OS),Windows_NT)
	SUFFIX := .exe
endif

ifndef VERBOSE
.SILENT: # no need for @
endif

PACKAGE := github.com/develerik/git-credential-1password
BUILD_HASH := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell git log -n1 --pretty='format:%cd' --date=format:'%Y-%m-%dT%H:%M:%S%z')
VERSION := $(VERSION)

# strip symbols
GO_BUILD_FLAGS := -s -w -extldflags "-static"
# define version
GO_BUILD_FLAGS := $(GO_BUILD_FLAGS) -X $(PACKAGE)/cmd.Version=$(VERSION)
# define build hash
GO_BUILD_FLAGS := $(GO_BUILD_FLAGS) -X $(PACKAGE)/cmd.Build=$(BUILD_HASH)
# define build date
GO_BUILD_FLAGS := $(GO_BUILD_FLAGS) -X $(PACKAGE)/cmd.Date=$(BUILD_DATE)
# go build command
GO_BUILD := CGO_ENABLED=0 go build
# go binary path
GO_BINPATH := $(shell go env GOPATH)/bin
# go linter
GO_LINTER := $(GO_BINPATH)/golangci-lint$(SUFFIX)
# go source files
GO_FILES = $(shell find . -name '*.go')

.DEFAULT_GOAL:=help

$(GO_LINTER):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_BINPATH) v1.31.0

##@ Build

.PHONY: git-credential-1password

git-credential-1password: bin/git-credential-1password$(SUFFIX) ## Build git-credential-1password

bin/git-credential-1password$(SUFFIX): $(GO_FILES)
	$(GO_BUILD) -ldflags '$(GO_BUILD_FLAGS)' -o bin/git-credential-1password$(SUFFIX) $(PACKAGE)

##@ Code Style

.PHONY: lint lint-super

lint: $(GO_LINTER) ## Lint the go source files
	$(GO_LINTER) run -c .github/linters/.golangci.yml

lint-super: ## Lint the whole project using GitHub's super-linter
	docker run -e RUN_LOCAL=true -v $(PWD):/tmp/lint github/super-linter

##@ Cleaning

.PHONY: clean

clean: ## Clean previous build
	rm -rf bin

##@ Helpers

.PHONY: help

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <command> \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
