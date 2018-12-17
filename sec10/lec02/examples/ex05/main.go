package main

import (
	"fmt"
	"io"
	"sync"
)

var (
	data = []string{"hello world!", "this is very nice", "how cool!"}
)

func main() {
	var wg sync.WaitGroup
	r, w := io.Pipe()
	producer(&wg, w)
	consumer(&wg, r)
	wg.Wait()
}

// consumer reads data from in
func consumer(wg *sync.WaitGroup, in io.Reader) {
	wg.Add(1)
	go func() {
		var err error
		var n int
		s := make([]byte, 256)
		for err == nil {
			n, err = in.Read(s)
			if err == nil || (err == io.EOF && n > 0) {
				fmt.Println(string(s[:n]))
			}
		}
		wg.Done()
	}()
}

// producer writes some strings to out
func producer(wg *sync.WaitGroup, out io.WriteCloser) {
	wg.Add(1)
	go func() {
		for _, s := range data {
			io.WriteString(out, s)
		}
		out.Close()
		wg.Done()
	}()
}
