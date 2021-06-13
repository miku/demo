SHELL := /bin/bash
GOFILES := $(shell find . -name "*.go")

main: $(GOFILES) wire
	go build -tags '!wireinject' -o $@

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
	rm -f main wire_gen.go

