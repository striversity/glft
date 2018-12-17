# Section 09 - Exercise 03 : Method Set

Complete the package 'wc' such that the golang application runs correctly. The package 'wc' provides a type WriteCounter, which implements the Writer interface shown below and the fmt.Stringer interface. WriteCounter must keep track of how many times the .Write() method is called on its value. *NOTE*: The value doesn't record or save the bytes.

## TODO 1 - Implement Writer interface

Given the Writer interface below, a value of WriteCounter records the number of times the `Write()` method was called. Hence, the return values of `Write()` is always `len(b), nil`.

```go
// Writer interface declares the Write method for sending data to a value
type Writer interface {
    // Write takes a buffer b of type []bytes and returns the number of bytes accepted.
    // If n < len(b), then Write will return an appropriate error explaining why.
    Write(b []byte) (n int, err error)
}
```

## TODO 2 - Implement fmt.Stringe interface

A value of type WriteCounter has the method `String() string` in its method set. When called, the method should return the number of times the `Write()` method was called. For example,
> There were 20 write operations