// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"log"
)

// Injectors from wire.go:

func initializeBar() string {
	myFooer := provideMyFooer()
	string2 := provideBar(myFooer)
	return string2
}

// wire.go:

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

	return f.Foo()
}

var Set = wire.NewSet(
	provideMyFooer, wire.Bind(new(Fooer), new(*MyFooer)), provideBar)

func main() {
	bar := initializeBar()
	log.Println(bar)
}
