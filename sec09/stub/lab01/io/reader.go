package io

type (
	// Reader interface declares the Read method for sending data to a value
	Reader interface {
		// Read takes a buffer 'b' of type []bytes and returns the number of bytes copied to 'b'.
		// If n < len(b), then Write will return an appropriate error explaining why.
		Read(b []byte) (n int, err error)
	}
)
