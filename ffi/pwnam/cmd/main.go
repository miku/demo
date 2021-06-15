package main

import (
	"fmt"

	"github.com/miku/demo/ffi/pwnam"
)

func main() {
	v := pwnam.Getpwnam("tir")
	fmt.Printf("user %d uses %s\n", v.Uid, v.Shell)
}
