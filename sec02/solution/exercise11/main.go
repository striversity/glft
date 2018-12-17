// TODO 4 - write a program which prints the agent's code name and full name
package main

import (
	"fmt"
	
	"github.com/striversity/glft/sec02/solution/exercise10/sa"
)

func main() {
	fmt.Printf("Agent Code Name: %v\n", sa.CodeName())
	fmt.Printf("Agent Full Name: %v\n", sa.FullName())
	// TODO 5 - demonstrates that private members of 'sa' is not exposed to the 'main' package
	// uncomment following lines to run this program
	fmt.Printf("Agent Salary: $%.2f", sa.salary)
}
