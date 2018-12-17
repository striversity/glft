// Section 08 - Lecture 03 : Using Pointers
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
	// implicit and explicit pointer deference
	// ----
	p := &Student{name: "John Doe", age: 45, ssn: "000-11-2222"}
	fmt.Printf("[DEBUG-1] p's value: %v, type: %T\n", p, p)
	fmt.Printf("[DEBUG-2] *p's value: %v, type: %T\n", *p, *p)

	student := *p
	fmt.Printf("[DEBUG-3] student's value: %v, type: %T\n", student, student)
	fmt.Printf("[DEBUG-4] Student Name (student.name): %v\n", student.name)
	fmt.Printf("Student Name ((*p).name): %v\n", *p.name)             // incorrect, 'Student.name' IS NOT a pointer
	fmt.Printf("[DEBUG-5] Student Name ((*p).name): %v\n", (*p).name) // explicit pointer deference
	fmt.Printf("[DEBUG-6] Student SSN (p.ssn): %v\n", p.ssn)          //  implicit pointer deference

}
