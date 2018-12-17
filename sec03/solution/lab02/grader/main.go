/*
# Section 03 - Lab 01 : Grader

## TODO 1

### Write a Go program which meets the following requirements

    1. Prompt the user for the number of test scores they would like to enter.
        - NOTE: value must be >= 0
    2. Prompts the user to enter test score.
    3. Repeat step 2 until all of the test scores have been entered.
    4. Compute letter grade for scores as follows:
		- A+ >= 98, A >= 95, B+ >= 90, B >= 80,  C+ >= 70, C >= 60, D+ >= 55, D >= 45, F < 45
*/
package main

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)

type (
	studentGrades []int
)

func main() {
	// prompt user for the number of test scores
	var numScores int
	prompt := "How many score will be entered: "
	for numScores = input.PromptInt(prompt); numScores < 0; {
		fmt.Printf("An invalid number was entered, %v. Please try again with a number >= 0.\n", numScores)
		numScores = input.PromptInt(prompt)
	}

	// allocate slice for scores
	testScores := make(studentGrades, numScores)
	prompt = "Please enter test score %v: "
	for i := range testScores {
		testScores[i] = input.PromptInt(fmt.Sprintf(prompt, i+1))
	}
	// calculate letter grade
	g := avgGrade(testScores)
	a := letterGrade(g)
	fmt.Printf("Your letter grade is %v for the scores %v\n", a, g)
}
func letterGrade(g int) string {
	// A+ >= 98, A >= 95, B+ >= 90, B >= 80,  C+ >= 70, C >= 60, D+ >= 55, D >= 45, F < 45
	letterGradeMap := []string{"A+", "A", "B+", "B", "C+", "C", "D+", "D"}
	numbericGradeMap := []int{98, 95, 90, 80, 70, 60, 55, 45}

	for i,v := range numbericGradeMap {
		if g >= v {
			return letterGradeMap[i]
		}
	}
	return "F"
}
func avgGrade(g studentGrades) int {
	var sum int
	for _, v := range g {
		sum += v
	}
	return sum / len(g)
}
