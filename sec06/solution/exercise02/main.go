// Section 06 - Exercise 02 : Optimal Timing for Multiple Prime Searchers Running Concurrently

package main

import (
	"errors"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
)

const (
	dataSize    = 100000 // 100,000
	numSearhers = 6
)

var (
	wg sync.WaitGroup
)

func main() {
	data, err := generateData()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numSearhers; i++ {
		d := subSlice(data, numSearhers, i)
		wg.Add(1)
		go primeSearcer(i+1, d)
	}

	wg.Wait()
}

// primeSearch checks if each number in the data set 'data' is a prime number. If it is, it prints it.
func primeSearcer(id int, data []int) {
	var i int
	for _, v := range data {
		for i = 2; i < v; i++ {
			if (v % i) == 0 {
				break // not a prime, try next int in data set
			}
		}
		// if i == v, this is a prime
		if i == v {
			fmt.Println(v)
		}
	}
	wg.Done()
}

// generateData produces a slice of 'dataSize' ints between 1 and maxInt
func generateData() (out []int, err error) {
	out = make([]int, dataSize)
	if out == nil {
		err = errors.New("Can't allocate memory for requested storage")
		return
	}

	for i := 0; i < dataSize; i++ {
		out[i] = i + 1
	}
	return
}
