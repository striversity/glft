package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/striversity/glft/sec09/stub/exercise03/wc"
)

type (
	// Writer interface declares the Write method for sending data to a value
	Writer interface {
		// Write takes a buffer b of type []bytes and returns the number of bytes accepted.
		// If n < len(b), then Write will return an appropriate error explaining why.
		Write(b []byte) (n int, err error)
	}
)

func main() {
	wc1 := new(wc.WriteCounter)
	var w Writer = wc1
	var buf []byte
	
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Int()%100; i++ {
		w.Write(buf)
	}
	fmt.Println(wc1)
}
