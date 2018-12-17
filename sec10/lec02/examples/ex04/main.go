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
		s := make([]byte, 256)
		for err == nil {
			_, err = in.Read(s)
			fmt.Println(string(s))
		}
		wg.Done()
	}()
}

// producer writes some strings to out
func producer(wg *sync.WaitGroup, out io.Writer) {
	wg.Add(1)
	go func() {
		for _, s := range data {
			io.WriteString(out, s)
		}
		wg.Done()
	}()
}
