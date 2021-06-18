// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/miku/demo/di/basic/extra"
	"log"
)

// Injectors from main.go:

// initializeBar creates a *Bar value given providers (constructor) in a set.
func initializeBar() *extra.Bar {
	logger := extra.NewLogger()
	foo := extra.NewFoo(logger)
	bar := extra.NewBar(foo)
	return bar
}

// main.go:

func main() {
	bar := initializeBar()
	log.Printf("bar: %#v", bar)
	log.Println("DI example")
}
