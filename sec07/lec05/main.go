// Section 07 - Lecture 05 : Channel Selection

package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	// inserting delays the wrong way
	// ----
	// fmt.Printf("Message 1 at: %v\n", time.Now())
	// sleep(1 * time.Second)
	// fmt.Printf("Message 2 at: %v\n", time.Now())

	// inserting delays with time.Sleep()
	// ----
	// fmt.Printf("Message 1 at: %v\n", time.Now())
	// time.Sleep(1 * time.Second)
	// fmt.Printf("Message 2 at: %v\n", time.Now())

	// get a notification form a channel
	// ----
	// fmt.Printf("Message 1 at: %v\n", time.Now())
	// alarm := notifyAfter(1 * time.Second)
	// <-alarm
	// fmt.Printf("Message 2 at: %v\n", time.Now())

	// get a notification from a channel with time.After()
	// ----
	// fmt.Printf("Message 1 at: %v\n", time.Now())
	// <-time.After(1 * time.Second)
	// fmt.Printf("Message 2 at: %v\n", time.Now())

	// the 'select' statement
	// ----
	// var ch1 chan int
	// select {
	// case <-ch1:
	// 	log.Info("Got data from ch1")
	// default:
	// 	log.Warn("No data from ch1")
	// }

	// selecting from multiple channels
	// ----
	// var ch1, ch2 chan string
	// select {
	// case <-ch1:
	// 	log.Info("Got data *FROM* ch1")
	// case ch2 <- "hello":
	// 	log.Info("Sent data *TO* ch2")
	// default:
	// 	log.Warn("No communication for ch1 or ch2")
	// }

	// generating random stream of bits
	// ----
	// fmt.Println("Random bit stream:")
	// const total = 10000
	// bits := randomBitsGen(total)
	// m:= make(map[int8]int)
	// for v := range bits {
	// 	m[v]++
	// }
	// for k,v := range m{
	// 	f:= (float32(v)/total)*100
	// 	fmt.Printf("%v occurred %.2f%% of the time\n", k, f)
	// }
	// fmt.Println()

	// timing out after waiting on a channel
	// ----
	d := producer()
	consumer(d)
}
func consumer(in chan string) {
	for {
		alarm := time.After(2 * time.Millisecond)
		select {
		case m, ok := <-in:
			if !ok {
				fmt.Println("No more data from 'in'")
				return
			}
			fmt.Printf("Consumer got - %v", m)
		case <-alarm:
			fmt.Println("Timedout waiting on for data")
			return
		}
	}
}
func producer() (out chan string) {
	out = make(chan string)
	go func() {
		count := 1
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer message %v\n", count)
			count++
		}
		time.Sleep(3 * time.Millisecond)
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer message %v\n", count)
			count++
		}
		close(out)
	}()
	return
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

func notifyAfter(delay time.Duration) (out chan time.Time) {
	out = make(chan time.Time)
	go func() {
		time.Sleep(delay)
		out <- time.Now()
		close(out)
	}()
	return out
}
func sleep(delay time.Duration) {
	end := time.Now().Add(delay)
	for time.Now().Before(end) {
		log.Info("Just waiting here, nothing to do...")
	}
}
