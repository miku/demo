package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file := createFileWithLines(
		"foo",
		"bar",
		"baz",
	)

	rr := &ReverseReader{file: file}
	// Seek to the end of file
	rr.SeekEnd()

	// Now we can use it like any other io.Reader, but sure it will read from tail to head
	scanner := bufio.NewScanner(rr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	check(scanner.Err())
}

type ReverseReader struct {
	file *os.File
}

// Seek to the final byte of the file
func (r *ReverseReader) SeekEnd() {
	_, err := r.file.Seek(0, io.SeekEnd)
	check(err)
}

// Read the file backwards
func (r *ReverseReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	// This no-op gives us the current offset value
	offset, err := r.file.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	var m int
	for i := 0; i < len(b); i++ {
		if offset == 0 {
			return m, io.EOF
		}
		// Seek in case someone else is relying on seek too
		offset, err = r.file.Seek(-1, io.SeekCurrent)
		if err != nil {
			return m, err // Should never happen
		}

		// Just read one byte at a time
		n, err := r.file.ReadAt(b[i:i+1], offset)
		if err != nil {
			return m + n, err // Should never happen
		}
		m += n
	}
	return m, nil
}

func createFileWithLines(data ...string) *os.File {
	tmp, err := ioutil.TempFile("", "")
	check(err)

	_, err = tmp.Write([]byte(strings.Join(data, "\n")))
	check(err)
	tmp.Close()

	file, err := os.Open(tmp.Name())
	check(err)

	return file
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
