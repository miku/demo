SHELL := /bin/bash

main: main.go hello.c
	go build -o $@ .

.PHONY: run
run:
	go run .

.PHONY: clean
clean:
	rm -f main
