package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	flag.Parse()
	var r io.Reader = os.Stdin
	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.Seek(0, io.SeekEnd); err != nil {
			log.Fatal(err)
		}
		r = f
	}
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		if err != nil {
			log.Printf("read failed: %v", err)
			break
		}
		fmt.Printf(line)
	}
}
