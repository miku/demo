//+build wireinject

package main

import (
	"log"

	"github.com/google/wire"
)

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

func provideBar(f Fooer) string {
	// f will be a *MyFooer.
	return f.Foo()
}

var Set = wire.NewSet(
	provideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	provideBar)

func initializeBar() (s string) {
	wire.Build(Set)
	return s
}

func main() {
	bar := initializeBar()
	log.Println(bar)
}
