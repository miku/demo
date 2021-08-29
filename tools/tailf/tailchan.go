package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	flag.Parse()
	var (
		r     = os.Stdin
		lines = make(chan string)
		wg    sync.WaitGroup
	)
	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
	reader := func() {
		defer wg.Done()
		br := bufio.NewReader(r)
		for {
			line, err := br.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("read failed: %v", err)
				break
			}
			lines <- line
		}
		close(lines)
	}
	writer := func() {
		defer wg.Done()
		for line := range lines {
			fmt.Printf(line)
		}
	}
	wg.Add(2)
	go writer()
	go reader()
	wg.Wait()
}
