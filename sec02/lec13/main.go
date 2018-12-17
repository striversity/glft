// Section 02 - Lecture 13 - Chapter Review
package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const huge = 9302948902983904828902830492893482093849283402839482948290348293902942893482039482903482098990899089
const (
	mon = iota
	tue = iota
	wed = iota
	thu = iota
	fri = iota
	sat = iota
	sun = iota
)

func foo(x int) {

}

func main() {
	// using defer()
	defer fmt.Println("leaving main()")
	// constants
	fmt.Println(float64(huge))
	fmt.Printf("mon: %v, tue: %v, wed: %v, thu: %v, fri: %v, sat: %v, sun: %v\n",
		mon, tue, wed, thu, fri, sat, sun)
	//variables
	_, _, y, _ := 9, 15, 21, 32
	fmt.Println(y)
	{ // scope-1
		x := 10
		{ // scope-2
			y := x * 5
			fmt.Println(x, y)
		} // 'y.scope-2' is no longer valid
		// but 'y from main is valid
	}

	// complex number
	c := 11 + 4i
	fmt.Printf("real(c): %v, imag(c): %v\n", real(c), imag(c))

	// blank idenfitier in 'if' and 'for'
	if i, _ := div(9, 15); i > 0 {
		fmt.Println("Division resulted in non-zero quoitent")
	}

	for _, i := div(9, 15); i > 0; i-- {
		fmt.Println("Number of items remaining")
	}
	// functions
	v()
	v(1)
	v(1, 5, 343, 98, 534, 593)

	if c, err := getItemCost(1); err == nil {
		fmt.Printf("Item 1 cost is $%.2f\n", c)
	}
	if _, err := getItemCost(0); err != nil {
		fmt.Println(err)
	}
	if _, err := getItemCost(-1); err != nil {
		fmt.Println(err)
	}

	hoo()
	f1 := retAfunc(2)
	f1()
	f2 := retAfunc(5)
	f2()

	log.Println("This is just some information")
	log.Println("some more information")
	log.Warnln("this is somewhat important")
	log.Errorln("this is REALLY  IMPORTANT")
}

func retAfunc(x int) func() {
	y := 15
	z := 100
	var voo = func() {
		fmt.Println("Hello, World!")
		fmt.Printf("x = %v, y = %v\n", x, y)
	}
	fmt.Println(z)
	return voo
}
func hoo() {
	defer fmt.Println("Leaving hoo()")
	defer fmt.Println("Entering hoo()")
}
func getItemCost(itemId int) (float64, error) {
	if itemId == 0 {
		return 0.0, errors.New("Invalid parameter, itemId must be > 0")
	}
	if itemId < 0 {
		return 0.0, errors.New("Negative values not allowed, itemId must be > 0")
	}
	return 12.94, nil
}
func v(f ...int) {
	fmt.Printf("f: %v, type: %T\n", f, f)
}
func div(x, y int) (int, int) {
	a := x / y
	b := x % y
	return a, b
}
