GO_BIN ?= go
all: install build
.PHONY: all

build:
	$(GO_BIN) build ./cmd/fdb

install:
	$(GO_BIN) install ./cmd/fdb
