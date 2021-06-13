SHELL := /bin/bash
GOFILES := $(shell find . -name "*.go")

main: $(GOFILES)
	go build -tags '!wireinject' -o $@

.PHONY: wire
wire:
	wire

.PHONY: ls-files
ls-files:
	@go list -f '{{.GoFiles}}'
	@go list -f '{{.GoFiles}}' -tags 'wireinject'
	@go list -f '{{.GoFiles}}' -tags '!wireinject'


.PHONY: clean
clean:
	rm -f main

