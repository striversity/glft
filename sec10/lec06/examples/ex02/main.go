package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Errorf() function
	// ----
	fmt.Println(foo())
	fmt.Println(goo())
}

func foo() error {
	return errors.New("This is an error from foo()")
}

func goo() error {
	rand.Seed(time.Now().UnixNano())
	return fmt.Errorf("This is an error from goo() - %v", rand.Int()%200)
}
