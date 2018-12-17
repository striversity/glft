package main

import (
	"fmt"
	"io"
	"os"

	"github.com/striversity/glft/sec10/lec01/ms"
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
	io.WriteString(rw, "Hello, world")
	io.WriteString(rw, ". ")
	io.WriteString(rw, "It is a wonderful day.")
	io.Copy(os.Stdout, rw)

	fmt.Println() // should print "Hello, world. It is a wonderful day."
}
