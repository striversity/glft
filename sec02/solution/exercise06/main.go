package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/* NOTE: I have provided the function 'getUserAge()',
	which gets the user's age. */
	userAge := getUserAge()
	fmt.Printf("You are %v yrs. old!\n", userAge)

	/* TODO 1 - complete the program using the value in the variable
	'userAge' by:
	 a) printing "You are too young to sign up." if 'userAge' is
	 less than or equal to 16,
	b) printing "Come back when you are 18.", if 'userAge' is equal
	to 17
	c) printing "Welcome, you are old enough to sign up!", if
	'userAge' is greater than or equal to 18
	*/
	switch {
	case userAge <= 16:
		fmt.Println("You are too young to sign up.")
	case userAge == 17:
		fmt.Println("Come back when you are 18.")
	default:
		fmt.Println("Welcome, you are old enough to sign up!")
	}

	// TODO 2 - run the program a few times to see different results
}

// ------- DON"T LOOK BELOW HERE
const (
	minAge = 1
	maxAge = 30
)

// getUserAge returns a value representing a user's age between minAge to maxAge
func getUserAge() (age uint8) {
	rand.Seed(time.Now().Unix())
	// get a random int value
	v := rand.Uint32()
	age = uint8(v%maxAge) + minAge
	return
}
