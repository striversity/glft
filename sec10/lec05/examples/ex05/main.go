package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxLines = 5
)

func main() {
	var buf []string
	line := 1
	fmt.Printf("Enter ATMOST %v lines of text below, (empty line to quit).\n", maxLines)
	fmt.Printf("#%v> ", line)
	// reading lines of text from a 'file'
	// ----
	scanner := bufio.NewScanner(os.Stdin)
	for line <= maxLines && scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}
		buf = append(buf, s)
		line++
		fmt.Printf("#%v> ", line)
	}
	// print text entered with Title case per line
	for i := range buf {
		s := strings.Title(buf[i])
		fmt.Printf("#%v: %v\n", i+1, s)
	}

}
