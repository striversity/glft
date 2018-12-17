/* TODO 1 - Write a complete Go program which declares the 
following identifiers as either const or var:

++ NOTE ++: Depending on the context, some of these might be 
const or var. So there isn't a single correct answer.

    pi - Used instead of the value 3.14
    secondsInHour - Used instead of the value 60
    hoursInDay - Used instead of the value 24
    presenterName - name of current presenter which may change as records are processed
    favoriteLanguage - holds your favoritate programming or written language
    itemCount - represents the number of items in a shopping cart
    totalPrice - price of all items in a shopping cart
    isLoggingEnabled
*/ 
package main

import "fmt"

/*
Taken from https://golang.org/ref/spec#Numeric_types
--------------------------- Numberic types -------------------
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

*/

const (
	pi               float32 = 2.14
	secondsInHour    uint8   = 60
	hoursInDay       uint8   = 24
	favoriteLanguage         = "golang"
	isLoggingEnabled         = true // if you plan to change logging dynamcially, then this would be 'var'
)

var (
	presenterName string
	itemCount     uint    // itemCount should never be a negative values, hence unsigned-int
	totalPrice    float64 // prepared to handle very large sums :), but float32 is fine too
)

func main() {
}
