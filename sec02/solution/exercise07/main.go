/* IMPORTANT: You *must* use 'continue' in at least one of 
your solutions.
---------------------------------------------------------------
TODO 1 - calculate the sum of ODD nubmers between 1 to 10001
HINT: Use the modulo operator: 
	https://www.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/what-is-modular-arithmetic

ODD numbers  -> 1, 3, 5, 7, 9, 11, .... 9999, 10001
Result should be: 25010001

TODO 2 - calculate the 'integer' average of all nubmers from 1 to 10001,
excluding the nubmers: 
   10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910
Result should be: 5003
*/

package main

import (
	"fmt"
)

const max = 10001

func main() {
	// TODO 1 - calculate the sum of ODD nubmers between 1 to 10001
	var sumOfOdd int
	for x := 1; x <= max; x++ {
		if x%2 == 1 {
			sumOfOdd += x
		} 
	}

	fmt.Printf("Sum of ODD nubmers from 1 to %v = %v\n", max, sumOfOdd)

	// TODO 2 - calculate the 'integer' average of all nubmers from 1 to 10001, excluding the nubmers:
	// 	10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910
	var count, sum int
	for x := 1; x <= max; x++ {
		switch x + 1 {
		case 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910:
			continue
		}
		count++
		sum += x
	}
	fmt.Printf("Avergae of 1 to %v, excluding 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910 is equal to %v\n",
		max, sum/count)
}
