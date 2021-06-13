SHELL := /bin/bash
GOFILES := $(shell find . -name "*.go")

main: $(GOFILES)
	go build -tags '!wireinject' -o $@,
// https://github.com/google/wire/issues/117#issue-407026505
.PHONY: wire
wire:
	wire

.PHONY: ls-files
ls-files:
	go list -f '{{.GoFiles}} {{.IgnoredGoFiles}}'
	go list -f '{{.GoFiles}} {{.IgnoredGoFiles}}' -tags 'wireinject'
	go list -f '{{.GoFiles}} {{.IgnoredGoFiles}}' -tags '!wireinject'


.PHONY: clean
clean:
	rm -f main

