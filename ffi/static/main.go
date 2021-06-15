package main

/*

int i = 0xFF;
static int j = 0x80;

*/
import "C"

import "fmt"

func main() {
	fmt.Printf("%d\n", C.i)
	// fmt.Printf("%d\n", C.j) // undefined reference to `j'
}
