# The binary to build (just the basename).
BIN := teltonika_exporter

BASEIMAGE ?= scratch
VERSION ?= dev
BINPATH ?= docker/$(BIN)
OS_NAME ?= $(shell uname -s | tr A-Z a-z)

.ONESHELL:
os:
	@echo bulding on $(OS_NAME)

build: os
	GOARCH=amd64 CGO_ENABLED=0 GOOS=$(OS_NAME) go build -v -ldflags "\
				-X main.version=$(VERSION) \
				" -o $(BINPATH)

run: build
	@cd docker
	@./$(BIN)