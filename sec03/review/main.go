// Section 03 - Lecture 06 : Review
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/verrol/glft/shared/input"
)

type (
	Row   []int
	Table []Row
)

func main() {
	// program arguments
	fmt.Printf("os.Args type: %T\n", os.Args)
	for i, v := range os.Args {
		fmt.Printf("os.Args[%v] = %v\n", i, v)
	}

	// convert string to numberic or boolean
	myBool, err := strconv.ParseBool(os.Args[3])
	fmt.Printf("err: %v\n", err)
	fmt.Printf("myBool: %v as type %T\n", myBool, myBool)
	myFloat, err := strconv.ParseFloat(os.Args[2], 32)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("myFloat: %.4v as type %T\n", myFloat, myFloat)

	// multiple-dimension arrays

	var table1, _ = initTable(3, 7, 4, 5)
	printTable(table1)
	table2, _ := initTable(4, 5, 2, 5, 3)
	printTable(table2)
}
func printTable(t Table) {
	fmt.Println("-----------------------------")
	for i := range t {
		for j := range t[i] {
			fmt.Printf("%5v", t[i][j])
		}
		fmt.Println()
	}

	// reslice a slice
	nums := [10]int{12, 53, 86, 94, 75, 10, 2, 15, 37, 40}
	s1 := nums[:5]
	s2 := s1[6:8]
	fmt.Printf("s1: %v, len(s1): %v\n", s1, len(s1))
	fmt.Printf("s2: %v, len(s2): %v\n", s2, len(s2))

	// revisiting 'string' iteration
	s := "Hello, 世界"
	l := len(s)
	fmt.Printf("s: %v, len(s): %v\n", s, l)
	for i := 0; i < l; i++ {
		fmt.Printf("s[%v] = %v, type: %T\n", i, s[i], s[i])
	}
	for i, v := range s {
		fmt.Printf("s[%v] = %c, type: %T\n", i, v, v)
	}
	// a string as a []byte
	sab := []byte(s)
	sar := []rune(s)
	fmt.Printf("sab: %v, len(sab): %v\n", sab, len(sab))
	fmt.Printf("sar: %v, len(sar): %v\n", sar, len(sar))
	var st2 = string(sab)
	fmt.Printf("st2: %v, len(st2): %v\n", st2, len(st2))
	var st3 = string(sar)
	fmt.Printf("st3: %v, len(st3): %v\n", st3, len(st3))

}
func initTable(rows int, cols ...int) (t Table, err error) {
	if rows != len(cols) {
		return nil, errors.New("Invalid number of cols specified")
	}
	t = make([]Row, rows)
	for i := range t {
		t[i] = make([]int, cols[i])
		for j := range t[i] {
			t[i][j] = input.GetRandInt(-20, 120)
		}
	}
	return t, nil
}
