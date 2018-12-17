// Section 11 - Lecture 04 : Generic TCP/IP Server & Client
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	Message struct {
		Source  string
		Payload []int
	}
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	// return to basics - producer and consumer using channels
	// ----
	wg := new(sync.WaitGroup)
	ch := producer(wg, "Producer 1")
	consumer(wg, ch)
	wg.Wait()
}

// consumer is a client that consume messages from server
func consumer(wg *sync.WaitGroup, in <-chan Message) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		timeout := time.After(5 * time.Second)
		for {
			select {
			case msg := <-in:
				processMessage(msg)
				timeout = time.After(5 * time.Second) // reset timer after we receive work
			case t := <-timeout:
				log.Warnf("Waited %v to send message, shutting down producer", t)
				return
			}
		}
	}()
}

// producer is a server the provides messages for clients
func producer(wg *sync.WaitGroup, src string) (out <-chan Message) {
	wg.Add(1)
	ch := make(chan Message)
	out = ch
	go func() {
		defer wg.Done()
		timeout := time.After(5 * time.Second)
		for {
			select {
			case ch <- newMessage(src):
				time.Sleep(time.Duration(r.Intn(2000)) * time.Millisecond) // random delay up to 2sec
				timeout = time.After(5 * time.Second)                      // reset timer after sending
			case t := <-timeout:
				log.Warnf("Waited %v to send message, shutting down producer", t)
				close(ch)
				return
			}
		}
	}()
	return
}

func newMessage(src string) Message {
	count := r.Intn(10) + 1 // 1 to 10
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = r.Intn(100)
	}
	return Message{Source: src, Payload: data}
}

func processMessage(msg Message) {
	total := 0

	for v := range msg.Payload {
		total += v
	}
	fmt.Printf("%v send:\n", msg.Source)
	l := len(msg.Payload)
	fmt.Printf("\tcount: %v, total: %v, avg: %v\n", l, total, float64(total)/float64(l))
}
