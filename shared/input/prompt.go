package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PromptBool gets a bool value from standard-in and retuns it
func PromptBool(prompt string) bool {
	input := getInputWithPrompt(prompt)
	b, err := strconv.ParseBool(input)

	for nil != err {
		fmt.Println(`Sorry, expected a boolean value.
			For example: 1, t, true, True, 0, f, false, False
			Please try again.`)
		input := getInputWithPrompt(prompt)
		b, err = strconv.ParseBool(input)
	}

	return b
}

// PromptFloat64 gets a floating point number from standard-in and retuns it
func PromptFloat64(prompt string) float64 {
	input := getInputWithPrompt(prompt)
	f, err := strconv.ParseFloat(input, 64)

	for nil != err {
		fmt.Println(`Sorry, expected an integer value. 
			For example: 3.14 or 1e6. 
			Please try again.`)
		input := getInputWithPrompt(prompt)
		f, err = strconv.ParseFloat(input, 64)
	}

	return float64(f)
}

// PromptInt gets an integer from standard-in and retuns it
func PromptInt(prompt string) int {
	input := getInputWithPrompt(prompt)
	i, err := strconv.ParseInt(input, 10, 32)

	for nil != err {
		fmt.Println("Sorry, expected an integer value. For example, 5. Please try again.")
		input := getInputWithPrompt(prompt)
		i, err = strconv.ParseInt(input, 10, 32)
	}

	return int(i)
}

// PromptString retuns the string read from standard-in, excluding
func PromptString(prompt string) string {
	input := getInputWithPrompt(prompt)
	return strings.TrimSpace(input)
}
func getInputWithPrompt(prompt string) (input string) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	for nil != err {
		fmt.Print(prompt)
		fmt.Println("Sorry, unable to read entry. Please try again.")
		input, err = reader.ReadString('\n')
	}
	input = input[:len(input)-1] // remove '\n' character
	return
}
