package ms

// TODO 4 - implement the method Read(b []byte) (n int, err error)
// TIP: Use copy() built-in function

// Read stores at most len(b) bytes in b from the memory store. Otherwise, it stores
// up to the available bytes and return how many bytes were stored. If n < len(b), then
// it returns ms.ErrEmpty
