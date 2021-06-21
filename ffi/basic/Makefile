SHELL := /bin/bash

main: main.go hello.c
	go build -o $@ .

.PHONY: run
run:
	go run .

_obj: main.go hello.c
	go tool cgo main.go

.PHONY: clean
clean:
	rm -rf main _obj
