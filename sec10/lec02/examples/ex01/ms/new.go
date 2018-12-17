package ms

import "errors"

// TODO 2 - implement a function that will properly initialize a MemStore and return it

// NewMemStore allocates storage of size 'cap' bytes and returns a properly constructed object.
func NewMemStore(cap uint) (store *MemStore, err error) {
	store = &MemStore{cap: cap}
	store.data = make([]byte, cap)
	if store.data == nil {
		return nil, errors.New("Unable to allocate requested storage")
	}

	return
}
