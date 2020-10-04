PWD := $(shell pwd)
GOBIN := $(PWD)/bin
REPOPATH := $(PWD)
GOFILES ?= $(shell find . -type f -name '*.go' -not -path "./pkg/*")
setup:
	GOBIN=$(GOBIN) GO111MODULE=on go get golang.org/x/tools/cmd/goimports
	GOBIN=$(GOBIN) GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.29.0

build:
	@echo "building IPWhitelist"
	@export REPOPATH=$(REPOPATH) && cd cmd/IPWServer && go run *.go

test:
	@echo "running IPWhitelist tests"
	@export REPOPATH=$(REPOPATH) && go test -v -cover -race ./...

ensure:
	@echo "downloading go dependencies"
	@ go mod download

fmt: setup
	@echo "running linting"
	@$(GOBIN)/goimports -local="IPWhitelist/" -w $(GOFILES)
	@cd ./packages && GO111MODULE=on $(GOBIN)/golangci-lint run ./...