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
		if i%10 == 0 {
			time.Sleep(2 * time.Second)
		}
	}
}
