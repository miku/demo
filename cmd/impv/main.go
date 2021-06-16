package main

import (
	"fmt"

	"github.com/miku/demo/misc/v1"
	v2 "github.com/miku/demo/misc/v2"
)

func main() {
	fmt.Println(misc.Version)
	fmt.Println(v2.Version)
}
