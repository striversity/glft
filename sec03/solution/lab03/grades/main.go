/*
# Section 03 - Lab02 : Class Test Report

This program prints a report for a class test score for a term.

## TODO 1

### Expand Lab01

    1. Instead of prompting the user for input, pass the following a program arguments:
        - number of students in the class, must be greater than 0
        - number of exams per term, must be greater than 0
        - maximum possible score per exam, must be between 100 and 120
		- class name, must not be an empty string
        * TIP: Use strconv.ParseInt() to convert string to number, eg: "45" to 45
	2. Use 30 a the minimum possible score
	3. Generate test score data using input.GetRandInt()

## TODO 2

### Print a report of the class's tests score and students grade

#### The report must include the following:

	1. Class name
	2. Student ID along with their test score for each test
	3. Student letter grade
	4. Average score for each test
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/shared/input"
)

type (
	// grades for one student
	StudentGrades []int
	// grades for students in a class
	ClassGrades []StudentGrades
)

const (
	// offset into os.Args
	programNameIdx = iota
	numStudentsIdx = iota
	numExamsIdx    = iota
	maxScoreIdx    = iota
	subjectNameIdx = iota
)

var (
	numStudents, numExams, maxScore int
	className                       string
)

func main() {
	// verify that we have enough parameters
	if 5 != len(os.Args) {
		log.Error("Invalid number of program arguments")
		printUsage()
		return // end program
	}

	var num int64
	var err error
	// validate nubmer of students per class is between 1 and 127
	if num, err = strconv.ParseInt(os.Args[numStudentsIdx], 10, 8); nil != err || num <= 0 {
		log.Errorf("Students per class must be between 1 and 127, got %v instead.", num)
		printUsage()
		return
	}
	numStudents = int(num)

	// validate number of exams is between 1 and 20
	if num, err = strconv.ParseInt(os.Args[numExamsIdx], 10, 8); nil != err || num <= 0 || num > 20 {
		log.Errorf("Number of exams must be between 1 and 20, got %v instead.", num)
		printUsage()
		return
	}
	numExams = int(num)

	// validate max score per test is between 100 and 120
	if num, err = strconv.ParseInt(os.Args[maxScoreIdx], 10, 8); nil != err || num < 100 || num > 120 {
		log.Errorf("Number of exams must be between 100 and 120, got %v intead.", num)
		printUsage()
		return
	}
	maxScore = int(num)

	className = os.Args[subjectNameIdx]
	// validate subject name is not empty
	if 0 == len(className) {
		log.Errorf("Class subject name no provided")
		printUsage()
		return
	}

	aClass := initGrades(numStudents, numExams, maxScore)
	printClassGrades(className, aClass)
}
func printClassGrades(cn string, grades ClassGrades) {
	// print report header
	fmt.Printf("Class: %v\n", strings.Title(cn))
	fmt.Printf("Students Grades\n")
	fmt.Println("---------------------------------------------------")
	fmt.Print("Student ID")
	f := "Test %v"
	for i := 0; i < len(grades[0]); i++ {
		s := fmt.Sprintf(f, i+1)
		fmt.Printf("%8v", s)
	}
	fmt.Printf("%14v\n", "Letter Grade")
	// print student test scores and term letter grade
	for i, s := range grades {
		fmt.Printf("%10v", i+1)
		for _, g := range s {
			fmt.Printf("%8v", g)
		}
		fmt.Printf("  %-v\n", letterGrade(avgGrade(s)))
	}
	// print average test score
	var testScores []int
	fmt.Printf("%-10v", "Avergage:")
	for i := 0; i < numExams; i++ {
		for j := 0; j < numStudents; j++ {
			testScores = append(testScores, grades[j][i])
		}
		fmt.Printf("%8v", avgGrade(testScores))
	}
	fmt.Println("\n---------------------------------------------------")
}
func initGrades(numStudents, numGrade, maxScore int) (c ClassGrades) {
	c = make(ClassGrades, numStudents)
	for j := range c {
		c[j] = make(StudentGrades, numGrade)
		for i := range c[j] {
			c[j][i] = input.GetRandInt(40, maxScore)
		}
	}
	return c
}
func letterGrade(g int) string {
	// A+ >= 98, A >= 95, B+ >= 90, B >= 80,  C+ >= 70, C >= 60, D+ >= 55, D >= 45, F < 45
	letterGradeMap := []string{"A+", "A", "B+", "B", "C+", "C", "D+", "D"}
	numbericGradeMap := []int{98, 95, 90, 80, 70, 60, 55, 45}

	for i, v := range numbericGradeMap {
		if g >= v {
			return letterGradeMap[i]
		}
	}

	return "F"
}
func avgGrade(g StudentGrades) int {
	var sum int
	for _, v := range g {
		sum += v
	}
	return sum / len(g)
}

// printUsage displays help on how to use the program
func printUsage() {
	fmt.Println("grades <numStudents> <numExams> <maxScore> <subject>")
	fmt.Println("\teg: grades 15 5 100 'Social Studies'")
	fmt.Println("\teg: grades 20 8 100 English")
}
