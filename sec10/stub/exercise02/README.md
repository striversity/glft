# Section 10 - Exercise 02 : Record writer

Complete the implemenation of a type RecordWriter, such that it implements io.Writer interface. The use-case is that a RecordWriter wraps it input with a record 'header' and 'footer'. The data between the 'header' and 'footer' is unchanged.

## TODO 1 - Declare type RecordWriter

## TODO 2 - Implement function NewRecordWriter

NewRecordWriter is a function that takes a io.Writer as input and return a RecordWriter value.

## TODO 3 - Implement io.Writer interface

When Write() is called on a RecordWriter value, it will prepend the string "[REC_HEADER_x]", followed by the given bytes, then finally it adds "[REC_FOOTER_x]". Where 'x' is ever increasing number for this RecordWriter.

## TODO 4 - Implement fmt.Stringer interface

A RecordWriter String() method return a string formated as: "x records written". For example,
> 20 records written

## REQUREMENTS

1. Be sure to test for nil receivers if you are using pointer receiver for your method(s)
2. Check for nil when allocation memory
3. Run the provided test functions by calling `go test` in the 'ms' package dir or individually in the test files.