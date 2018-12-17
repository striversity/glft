# Section 09 - Exercise 04 : interface and functions

Complete the package 'wc' such that the golang application runs correctly. The package 'wc' provides a type WriteCounter, which implements the Writer interface shown below and the fmt.Stringer interface. WriteCounter must keep track of how many times the .Write() method is called *AND* how many bytes were writen to the value. *NOTE*: The value doesn't record or save the bytes.

## TODO 1 - Define and appropriate type WriteCounter

Declare the type WriteCounter such that it can store both the number of writes and number of bytes provided to the value using the .Write() method. See TODO 2

## TODO 2 - Implement Writer interface

Given the Writer interface below, a value of WriteCounter records the number of times the `Write()` method was called. Hence, the return values of `Write()` is always `len(b), nil`.

```go
// Writer interface declares the Write method for sending data to a value
type Writer interface {
    // Write takes a buffer b of type []bytes and returns the number of bytes accepted.
    // If n < len(b), then Write will return an appropriate error explaining why.
    Write(b []byte) (n int, err error)
}
```

## TODO 3 - Implement fmt.Stringe interface

A value of type WriteCounter has the method `String() string` in its method set. When called, the method should return the number of times the `Write()` method was called. For example,
> There were 6 write operations totaling 664 bytes

## TODO 4 - implement the function storeData()

The function main.storeData() takes a Writer and []string, and writes each string in the slice to the Writer.
