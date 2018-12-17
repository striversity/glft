// Section 08 - Lecture 02 : Creating Pointers
package main

import "fmt"

type (
	ID      string
	Student struct {
		name string
		age  uint8
		ssn  ID
	}
)

func main() {
	// allowed pointers from literals
	// ----
	pStudent := &Student{}
	fmt.Printf("pStudent's value: %v, type: %T\n", pStudent, pStudent)
	// pointer to initialized Student literal
	pStudent = &Student{name: "John Doe", age: 45, ssn: "000-11-2222"}
	fmt.Printf("pStudent's value: %v, type: %T\n", pStudent, pStudent)

	// pointer to map literal
	pMap := &map[string]string{"name": "Jane Smith", "age": "35", "weight": "110lb"}
	fmt.Printf("pMap's value: %v, type: %T\n", pMap, pMap)
}
