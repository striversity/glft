package io

type (
	// Writer interface declares the Write method for sending data to a value
	Writer interface {
		// Write takes a buffer b of type []bytes and returns the number of bytes accepted.
		// If n < len(b), then Write will return an appropriate error explaining why.
		Write(b []byte) (n int, err error)
	}
)
