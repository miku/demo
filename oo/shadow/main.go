package main

import "fmt"

type A struct {
	name string
}

func (a A) Greet() { fmt.Printf("[A] Hello %s\n", a.name) }

type B struct {
	A
}

func (b B) Greet() {
	fmt.Printf("[B] Hi %s\n", b.name)
	b.A.Greet()
}

func main() {
	var a A
	var b B
	a.name = "Anna"
	b.name = "Bela"
	a.Greet()
	b.Greet()
}

// [A] Hello Anna
// [B] Hi Bela
// [A] Hello Bela
