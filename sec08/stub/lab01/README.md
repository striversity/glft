# Section 08 - Lab 01 : Playing with pointers

Compare array bytes vs slice of byte in pointer to struct.

## TODO - Update Document.Data to use []byte instead of [100e6]byte

Update the example application form Sec08 Lec04 Example 06, such that the Document structure using a []byte instead of an array of []. The Document.Data must still be initialized with an 100 MegaByte slice.

Compare the runtimes between the two implementations. Is there a significant performance differnce between the pointer to Document.Data using an array bytes vs a pointer to Document.Data using []byte?
