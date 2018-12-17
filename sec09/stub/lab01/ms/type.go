package ms

import "errors"

var (
	// ErrFull is returned when Write() can't store all the bytes
	ErrFull = errors.New("Storage capacity reached")
	// ErrEmpty is returned when Read() is called on an empty memory store
	ErrEmpty = errors.New("No more data available")
)

// TODO 1 - declare type MemStore
type (
	// MemStore is an in-memory data store which allows for storing and retriveing slice of bytes.
	// NOTE: User of MemStore must first specified the capacity
	// TIP:
	//                data buffer: [xxxxxxx-------]
	//                              ↑      ↑
	//                              |      |
	//  read data from this offset -+      +- write data from this offset
)
