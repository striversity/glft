package wc

import (
	"errors"
	"fmt"
)

// TODO 1 - declare type Writecounter
type (
	// WriteCounter keeps track of the following, number calls to Write() and the number of bytes written
	WriteCounter struct {
		numWrites uint
		numBytes  uint
	}
)

// TODO 2 - implement the method Write(b []byte) (n int, err error)
// Write updates the value each time this method is called and the number of bytes written
func (r *WriteCounter) Write(b []byte) (n int, err error) {
	if nil == r {
		return 0, errors.New("Invalid parameter, 'nil' value")
	}

	r.numWrites++
	r.numBytes += uint(len(b))
	return len(b), nil
}

// TODO 3 - imeplement the method String() string
// String return an informative message about how the Write() was used
func (r WriteCounter) String() string {
	return fmt.Sprintf("There were %v write operations totaling %v bytes", r.numWrites, r.numBytes)
}
