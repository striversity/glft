// Section 10 - Lecture 04 : Formatted Output
package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	Println(true)
	Println(2010)
	Println(9.15)
	Println(7 + 7i)
	Println("Hello world!")
	Println(ID("19520925"))
	Println(&Person{name: "Jane Doe"})
}
// Println is my simple println function
func Println(v ...interface{}) {
	Print(v...)
	Print("\n")
}
func Print(v ...interface{}) {
	for i, x := range v {
		switch x.(type) {
		case bool:
			fmt.Print(x.(bool))
		case int:
			fmt.Print(x.(int))
		case float64:
			fmt.Print(x.(float64))
		case complex128:
			fmt.Print(x.(complex128))
		case string:
			fmt.Print(x.(string))
		case *Person:
			fmt.Print(x.(*Person))
		default:
			fmt.Printf("type: %v.%T", x, x)
		}
		if i > 0 {
			fmt.Print(" ")
		}
	}
}

