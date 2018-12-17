package main

import "fmt"

func main() {
	// scan int from string
	// ----
	input := "41" // we could you strconv.ParseInt() too

	var a int
	fmt.Sscan(input, &a)
	fmt.Printf("The answer to the universe is: %v\n", a+1)
}
