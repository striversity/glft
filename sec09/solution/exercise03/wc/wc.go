package wc

import (
	"errors"
	"fmt"
)

type (
	// WriteCounter keeps track of how many times the method Write() was called
	WriteCounter uint 
)

// TODO 1 - implement the method Write(b []byte) (n int, err error)
// Write updates the value each time this method is called
func (r *WriteCounter) Write(b []byte) (n int, err error) {
	if nil == r {
		return 0, errors.New("Invalid parameter, 'nil' value")
	}

	*r += WriteCounter(1)
	return len(b), nil
}

// TODO 2 - imeplement the method String() string
// String return an informative message about how many times the method Write() was called
func (r WriteCounter) String() string {
	return fmt.Sprintf("There were %v write operations", uint(r))
}
