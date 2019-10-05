#!/bin/bash

export NOW=$(shell date +"%Y-%m-%d")

PACKAGE = github.com/reyhanfahlevi/simple-app
COMMIT_HASH = $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE = $(shell date +%FT%T%z)
ldflags = -X $(PACKAGE)/cmd.CommitHash=$(COMMIT_HASH) -X $(PACKAGE)/cmd.BuildDate=$(BUILD_DATE) -s -w

build:
	@echo "${NOW} == BUILDING HTTP SERVER API"
	@echo "${ldflags}"
	@CGO_ENABLED=0 go build -ldflags '$(ldflags)' -o http_api cmd/api/main.go