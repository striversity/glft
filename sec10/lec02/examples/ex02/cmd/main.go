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
	// n, err := io.WriteString(nil, "hello")
	// log.Warnf("n: %v, err: %v", n, err)
	WriteString(rw, "Hello, world")
	WriteString(rw, ". ")
	WriteString(rw, "It is a wonderful day.")
	Copy(os.Stdout, rw)

	fmt.Println() // should print "Hello, world. It is a wonderful day."
}
func Copy(dst io.Writer, src io.Reader) (n int64, err error) {
	p := make([]byte, 256)

	for err == nil {
		c, readErr := src.Read(p)
		if c > 0 {
			c, err = dst.Write(p[:c])
			n += int64(c)
		}

		if readErr == io.EOF {
			return
		}
		if readErr != nil {
			return n, readErr
		}
	}
	return
}

// WriteString converts a string to []bytes and then writes it to an io.Writer
func WriteString(w io.Writer, s string) (n int, err error) {
	if w == nil {
		return 0, io.ErrShortWrite
	}
	return w.Write([]byte(s))
}
