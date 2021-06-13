SHELL := /bin/bash

main: main.go wire
	go build -o $@ main.go

.PHONY: wire
wire:
	wire

.PHONY: clean
clean:
	rm -f main

