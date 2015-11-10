GO ?= go
GOPATH := $(CURDIR)/_vendor:$(GOPATH)

all: build

build:
	$(GO) build

install:
	$(GO) install

.PHONY: install-deps
install-deps:
	go get gopkg.in/yaml.v2
	go get github.com/codegangsta/cli

.PHONY: test
test:
	go test src/sh/command_test.go
