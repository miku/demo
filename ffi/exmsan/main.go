package main

// Test passing C struct to exported Go function.

/*
#include <stdint.h>

typedef struct { uintptr_t x, y; char b; } T;

void CF(int x);
*/
import "C"

//export F
func F(t C.T) { println(t.x, t.y, t.b) } // ./main.go:15:41: t.zzz undefined (type _Ctype_struct___0 has no field or method zzz)

func main() {
	C.CF(C.int(0))
	// C.CF(123) // ok
}
