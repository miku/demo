package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	for {
		fmt.Printf("line %d\n", i)
		i += 1
		time.Sleep(100 * time.Millisecond)
	}
}
