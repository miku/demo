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
	log.Printf("%#v", bar)
}
