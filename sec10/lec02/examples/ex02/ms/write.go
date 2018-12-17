package ms

import (
	"errors"
)

// TODO 3 - implement the method Write(b []byte) (n int, err error)
// Write stores len(b) bytes if there is space in the memory store. Otherwise, it stores
// up to the available space and return how many bytes were stored. If n < len(b), then
// it returns ms.ErrFull
func (m *MemStore) Write(b []byte) (n int, err error) {
	if nil == m {
		return 0, errors.New("Write called on nil value")
	}

	n = copy(m.data[m.writeOffset:], b)
	m.writeOffset += n
	if n < len(b) {
		return n, ErrFull
	}
	return n, nil
}
