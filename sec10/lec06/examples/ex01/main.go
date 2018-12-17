package main

import (
	"io"
	"os"
)

func main() {
	// os.Stderr - standard error 'file'
	// ----
	io.WriteString(os.Stdout, "This is to stdout\n")
	io.WriteString(os.Stderr, "This is to stderr\n")

}
