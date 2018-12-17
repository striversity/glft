// Section 07 - Lecture 05 : Channel Selection

package main

import (
	"fmt"
)

func main() {
	// generating random stream of bits
	// ----
	fmt.Println("Random bit stream:")
	const total = 10000
	bits := randomBitsGen(total)
	m := make(map[int8]int)
	for v := range bits {
		m[v]++
	}
	for k, v := range m {
		f := (float32(v) / total) * 100
		fmt.Printf("%v occurred %.2f%% of the time\n", k, f)
	}
	fmt.Println()

}

func randomBitsGen(l int) (out chan int8) {
	out = make(chan int8)

	go func() {
		for i := 0; i < l; i++ {
			select {
			case out <- 0:
			case out <- 1:
			}
		}
		close(out)
	}()

	return
}
