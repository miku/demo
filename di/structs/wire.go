//+build wireinject

package main

import (
	"log"

	"github.com/google/wire"
)

type Foo int
type Bar int

func ProvideFoo() Foo {
	return Foo(1)
}

func ProvideBar() Bar {
	return Bar(2)
}

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"))

func initializeFooBar() *FooBar {
	wire.Build(Set)
	return &FooBar{}
}

func main() {
	fooBar := initializeFooBar()
	log.Printf("%#v", fooBar)
}
