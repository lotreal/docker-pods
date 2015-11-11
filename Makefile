GO ?= go
GOPATH := $(CURDIR)/vendor:$(GOPATH)

all: build

build:
	$(GO) build

install:
	$(GO) install

.PHONY: install-deps
install-deps:
	$(GO) get gopkg.in/yaml.v2
	$(GO) get github.com/codegangsta/cli
	$(GO) get github.com/stretchr/testify

.PHONY: run
run:
	$(GO) run -v main.go

.PHONY: test
test:
	$(GO) test -v src/sh/command_test.go
