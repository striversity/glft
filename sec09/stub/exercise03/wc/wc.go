package wc

import (
	"errors"
	"fmt"
)

type (
	// WriteCounter keeps track of how many times the method Write() was called
	WriteCounter uint 
)

// TODO 1 - implement the method Write(b []byte) (n int, err error) for WriteCounter
// Write updates the value each time this method is called

// TODO 2 - imeplement the method String() string for WriteCounter
// String return an informative message about how many times the method Write() was called
