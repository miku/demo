//+build wireinject

// Note: The build tag here must come before the package declaration; cf.
// https://github.com/google/wire/issues/178#issuecomment-493633100,
// https://github.com/google/wire/issues/178#issuecomment-543349470,
// https://github.com/google/wire/issues/117#issue-407026505
package main

import (
	"log"

	"github.com/google/wire"
	"github.com/miku/demo/di/basic/extra"
)

// initializeBar creates a *Bar value given providers (constructor) in a set.
func initializeBar() *extra.Bar {
	wire.Build(extra.DefaultSet)
	return &extra.Bar{}
}

func main() {
	bar := initializeBar()
	log.Printf("bar: %#v", bar)
	log.Println("DI example")
}
