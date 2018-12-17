package ms

import (
	"errors"
	"io"
)

// TODO 4 - implement the method Read(b []byte) (n int, err error)
// Read stores at most len(b) bytes in b from the memory store. Otherwise, it stores
// up to the available bytes and return how many bytes were stored. If n < len(b), then
// it returns io.EOF
func (m *MemStore) Read(b []byte) (n int, err error) {
	if nil == m {
		return 0, errors.New("Read called on nil value")
	}

	n = copy(b, m.data[:m.writeOffset])        // copy out n bytes to 'b'
	r := copy(m.data, m.data[n:m.writeOffset]) // shift data to start of buffer, move data[n:writeOffset] -> data[0:writeOffset-n]
	m.writeOffset = r                          // bytes remaining in buffer
	if n < len(b) {
		return n, io.EOF
	}
	return n, nil
}
