package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	help     bool
	pwLength = 16
)

func init() {
	flag.BoolVar(&help, "h", false, "show this help ")
	flag.IntVar(&pwLength, "l", pwLength, "password length 8-255")
	flag.Parse()
}

func main() {
	if help || pwLength < 8 || pwLength > 255 {
		flag.Usage()
		return
	}

	pw := randomPwGen(pwLength)
	for v := range pw {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}

// randomPwGen creates a channel into which 'l' random characters are written.
// NOTE: Perfer more letters to numbers, and more numbers to special characters.
func randomPwGen(l int) (out chan rune) {
	const (
		letters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
		numbers = `0123456789`
		special = `-=~!@#$%^&*()_+[]\{}|;':",./<>?`
	)
	var (
		lettersRune = []rune(letters)
		numbersRune = []rune(numbers)
		specialRune = []rune(special)
		s           = rand.NewSource(time.Now().Unix())
		r           = rand.New(s)
	)
	out = make(chan rune, l)
	for {
		rn := r.Int()
		select {
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:

		case out <- numbersRune[rn%len(numbersRune)]:
		case out <- numbersRune[rn%len(numbersRune)]:
		case out <- numbersRune[rn%len(numbersRune)]:

		case out <- specialRune[rn%len(specialRune)]:
		default:
			close(out)
			return
		}
	}
}
