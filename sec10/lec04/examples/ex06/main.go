// Section 10 - Lecture 04 : Formatted Output
package main

import (
	"fmt"
)

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
	s := Sprint(v...)
	fmt.Print(s)
}
func Sprint(v ...interface{}) (s string) {
	for i, x := range v {
		switch x.(type) {
		case bool:
			if x.(bool) {
				s += "true"
			}
			s += "false"
		case int:
			s += intToStr(x.(int))
		case float64:
			s += f64ToStr(x.(float64))
		case complex128:
			s += cmplToStr(x.(complex128))
		case string:
			s += x.(string)
		case *Person:
			s += ppToStr(x.(*Person))
		default:
			s += unknownToStr(x)
		}
		if i > 0 {
			s += " "
		}
	}
	return
}

func intToStr(i int) string {
	return fmt.Sprint(i)
}
func f64ToStr(f float64) string {
	return fmt.Sprint(f)
}
func cmplToStr(c complex128) string {
	return fmt.Sprintf("(%v+%vi)", real(c), imag(c))
}
func ppToStr(p *Person) string {
	return fmt.Sprintf("%v", p)
}
func unknownToStr(v interface{}) string {
	return fmt.Sprintf("%#v.%T", v, v)

}
