// Dining Philosophers Problem
// Implementation of a golang abritratior solution to Dining Philosophers Problem.
// https://en.wikipedia.org/wiki/Dining_philosophers_problem
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type (
	Fork struct{}
	// Waiter keeps a set of forks for dinning philophers
	Waiter struct {
		forks chan *Fork
	}
)

const (
	numPhilosophers = 5
	dinnerDuration  = 10 * time.Second
)

var (
	s      = rand.NewSource(time.Now().Unix())
	r      = rand.New(s)
	waiter = &Waiter{
		forks: make(chan *Fork, numPhilosophers),
	}
	wg sync.WaitGroup
)

func main() {
	for i := 0; i < numPhilosophers; i++ {
		waiter.forks <- &Fork{}
		philosopher(i)
	}

	wg.Wait() // wait for all Philosophers to get up from the table :)
	fmt.Println("All Philosophers finished thinking and eating...")
}

// philosopher implements the two things a Philosopher does, think and eat
func philosopher(id int) {
	wg.Add(1)
	// a philosopher's work is to think for some time or eat for sometime
	go func() {
		defer wg.Done()
		endOfDinner := time.After(dinnerDuration) // how long to dine
		for {
			select {
			case <-endOfDinner:
				return
			default:
				if r.Int()%2 == 0 { // think or eat?
					think(id)
				} else {
					eat(id)
				}
			}
		}
	}()
}

func think(id int) {
	d := time.Duration(r.Int()%1000) * time.Millisecond // how long to think
	//fmt.Printf("Philosopher %v is thinking\n", id)
	time.Sleep(d)
}
func eat(id int) {
	fmt.Printf("Philosopher %v needs to eat\n", id)
	// get two forks to eat
	start := time.Now()
	fork1, fork2 := waiter.getForks()
	fmt.Printf("Philosopher %v waited %v to eat\n", id, time.Since(start))
	d := time.Duration(r.Int()%1000) * time.Millisecond // how long to eat
	time.Sleep(d)
	waiter.returnForks(fork1, fork2) // must return forks after eating
}

// getForks MUST return two forks when requested
func (w *Waiter) getForks() (fork1, fork2 *Fork) {
	for {
		select {
		case fork1 = <-w.forks:
			select {
			case fork2 = <-w.forks:
				return
			default:
				w.forks <- fork1
			}
		}
		// wait some time before trying to get two forks again
		time.Sleep(1 * time.Millisecond)
	}

	return
}
func (w *Waiter) returnForks(fork1, fork2 *Fork) {
	w.forks <- fork1
	w.forks <- fork2
}
