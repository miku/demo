package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("filename required")
	}
	filename := flag.Arg(0)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Seek(0, io.SeekEnd); err != nil {
		log.Fatal(err)
	}
	// setup watch
	fd, err := syscall.InotifyInit()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := syscall.InotifyAddWatch(fd, filename, syscall.IN_MODIFY); err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(f)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			wait(fd)
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf(line)
	}
}

func wait(fd int) error {
	for {
		var buf [syscall.SizeofInotifyEvent]byte
		if _, err := syscall.Read(fd, buf[:]); err != nil {
			return err
		}
		r := bytes.NewReader(buf[:])
		var ev = syscall.InotifyEvent{}
		if err := binary.Read(r, binary.LittleEndian, &ev); err != nil {
			return err
		}
		if ev.Mask&syscall.IN_MODIFY == syscall.IN_MODIFY {
			return nil
		}
	}
}
