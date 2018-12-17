# Section 09 - Lab 01 : Memory Store

Complete the implemenation of a type MemStore, such that it implements an in-memory byte storage. The use-case is that a MemStore value can accept a number bytes written to it using Write() method up to the specified capacity. Bytes can then be read from the MemStore using Read() method.

## TODO 1 - Declare type MemStore

## TODO 2 - Implement function NewMemStore

NewMemStore is a function that take the capacity for the memstore and return a valid memstore if storage can be allocated.

## TODO 3 - Implement io.Writer interface

io.Writer is defined in _peer_ directory 'io'. When called on the memstore, Write() method stores the specified bytes if capacity allows it. Otherwise, the method stores as much as possible and report that the store is _full_.

## TODO 4 - Implement io.Reader interface

io.Reader is defined in _peer_ directory 'io'. When called on the memstore, Read() method stores the specified bytes if avaliable in the store, into the specified slice parameter. Otherwise, the method stores as much as possible and report that the store is _empty_.

## REQUREMENTS

1. Be sure to test for nil receivers if you are using pointer receiver for your method(s)
2. Check for nil when allocation memory
3. Run the provided test functions by calling `go test` in the 'ms' package dir or individually in the test files.