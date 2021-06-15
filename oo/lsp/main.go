package main

// If you need substitution, use an interface.

type Reader interface {
	// Read reads up to len(buf) bytes into buf.
	Read(buf []byte) (n int, err error)
}

// For more implementations, see: github.com/miku/exploreio
