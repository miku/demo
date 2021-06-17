package main

import (
	"fmt"

	"github.com/miku/demo/localpkg/learn"
	"github.com/miku/demo/misc/v1"
	v2 "github.com/miku/demo/misc/v2"

	// If package name does not match directory, an alias is used.
	anyname "github.com/miku/demo/misc/namediff"
	other "github.com/miku/demo/misc/namediff/nested"
)

func main() {
	fmt.Println(misc.Version)
	fmt.Println(v2.Version)

	anyname.Render()
	fmt.Println(other.Nested())
	fmt.Println("learning go?", learn.Golang())
}
