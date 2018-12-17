package main

import (
	"fmt"

	"github.com/striversity/glft/sec09/stub/lab01/io"

	"github.com/striversity/glft/sec09/stub/lab01/ms"
)

type (
	// ReadWriter combines the interfaces io.Reader and io.Writer
	ReadWriter interface {
		io.Reader
		io.Writer
	}
)

func main() {
	var m, _ = ms.NewMemStore(100)
	test(m)
}

func test(rw ReadWriter) {
	var totalBytes int
	n, _ := rw.Write([]byte("Hello, world"))
	totalBytes += n

	n, _ = rw.Write([]byte(". "))
	totalBytes += n

	n, _ = rw.Write([]byte("It is a wonderful day."))
	totalBytes += n

	b := make([]byte, 1)
	_, err := rw.Read(b)
	for err == nil {
		fmt.Print(string(b))
		_, err = rw.Read(b)
	}
	fmt.Println() // should print "Hello, world. It is a wonderful day."
}
