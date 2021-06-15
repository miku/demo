package main

/*
#cgo CFLAGS: -g -Wall
#include "sum.h"
*/
import "C"

import "fmt"

func main() {
	s := C.sum(1, 1)
	fmt.Printf("%T %v\n", s, s) // main._Ctype_long 2

	// var v int64
	// v = s * 2
	// fmt.Printf("%T %v\n", v, v)
	// ./main.go:15:4: cannot use s * 2 (type _Ctype_long) as type int64 in assignment

	var (
		v  C.long = 2
		w  C.long = v * s
		gv int    = int(v)
		gw int64  = int64(w)
	)
	fmt.Printf("%T %d\n", w, w)
	fmt.Printf("%T %d\n", gv, gv)
	fmt.Printf("%T %d\n", gw, gw)
}
