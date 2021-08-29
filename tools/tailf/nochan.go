package main

import (
	"fmt"
	"strings"
)

type LineBuffer struct {
	writeHead int
	n         int
	lines     []string
}

func NewLineBuffer(n int) *LineBuffer {
	if n < 1 {
		panic("minimum buffer size is one")
	}
	return &LineBuffer{
		n:     n,
		lines: make([]string, n),
	}
}

func (b *LineBuffer) Push(s string) {
	b.lines[b.writeHead] = s
	b.writeHead += 1
	if b.writeHead == b.n {
		b.writeHead = 0
	}
}

func (b *LineBuffer) String() string {
	var sb strings.Builder
	for i := 0; i < b.n; i++ {
		if i == b.writeHead {
			fmt.Fprintf(&sb, "-> [%d] %s\n", i, b.lines[i])
		} else {
			fmt.Fprintf(&sb, "   [%d] %s\n", i, b.lines[i])
		}
	}
	return sb.String()
}

func main() {
	// flag.Parse()
	// if flag.NArg() == 0 {
	// 	log.Fatal("file required")
	// }
	// f, err := os.Open(flag.Arg(0))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// br := bufio.NewReader(f)

	buf := NewLineBuffer(3)
	buf.Push("hello")
	buf.Push("world")
	buf.Push("abc")
	buf.Push("def")
	fmt.Println(buf)
}
