SHELL := /bin/bash

main: main.go extra.c
	CC=clang go build -o $@ -msan .

.PHONY: run
run:
	go run .

.PHONY: clean
clean:
	rm -f main
