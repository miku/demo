package main

/*
#cgo CFLAGS: -g -Wall
#include "hello.h"
*/
import "C"

import "fmt"

func main() {
	p := C.hello()
	fmt.Println(C.GoString(p))
}
