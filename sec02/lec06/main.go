// Section 02 - Lecture 06 : The "switch" statement
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 'if-if else-else' statement
	if x := 11; x < 5 {
		fmt.Println("x is less than 5")
	} else if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is eitehr 6, 7, 8, 9, or 10")
	}

	// 'switch' statement
	switch x := 4; {
	case x < 5:
		fmt.Println("x is less than 5")
	case x > 10:
		fmt.Println("x is greater than 10")
	case x == 5:
		fmt.Println("x is equal to 5")
	default:
		fmt.Println("x is eitehr 6, 7, 8, 9, or 10")
	}

	// switch statement using strings
	switch day := getDay(); day {
	case "Monday":
		fmt.Printf("Yippie, it is %v, the first day of the week!\n", day)
	case "Tuesday", "Wednesday":
		fmt.Printf("Today is %v, time to get some work done.\n", day)
	case "Thursday":
		fmt.Printf("It is only %v, one more day to Friday.\n", day)
	case "Friday":
		fmt.Printf("TGIF, time for %v night movie!\n", day)
	case "Saturday":
		fmt.Println("Arrg, too many chores, but...")
		fallthrough
	case "Sunday":
		fmt.Printf("It is %v, and I am still enjoying the weekend.\n", day)
	default:
		fmt.Printf("Waht day is this? %v\n", day)
	}
}

// ------- DON"T LOOK BELOW HERE
var daysOfWeek = map[int]string{
	0: "Monday",
	1: "Tuesday",
	2: "Wednesday",
	3: "Thursday",
	4: "Friday",
	5: "Saturday",
	6: "Sunday",
}

// getDay returns a day of the week as a string
func getDay() (d string) {
	rand.Seed(time.Now().Unix())
	// get a random int value
	v := int(rand.Uint32())
	d = daysOfWeek[(v % len(daysOfWeek))]
	return
}
