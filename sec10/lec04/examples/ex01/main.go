// Section 10 - Lecture 04 : Formatted Output
package main

import (
	"io"
	"os"
)

var (
	data = []string{
		"Hello, goood morning!",
		"Hi, my name is Sam.",
		"Hey, good to see you again",
		"Welcome back",
	}
)

func main() {
	// writing to named file
	// ----
	out, _ := os.Create("data.txt")
	defer out.Close()

	for _, s := range data {
		io.WriteString(out, s)
		io.WriteString(out, "\n")
	}
}
