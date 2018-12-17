package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	maxMsgPerProducer = 100
	numProducers      = 3
	numConsumers      = 2
	s                 = rand.NewSource(time.Now().Unix())
	r                 = rand.New(s)
	ch                chan string
	wgProducers       sync.WaitGroup
	wgConsumers       sync.WaitGroup
	consumerMutex     sync.Mutex
)

type (
	ProdInfo struct {
		sum   int
		count int
	}
)

func init() {
	flag.IntVar(&maxMsgPerProducer, "m", maxMsgPerProducer, "max number of messages per producer, m > 0")
	flag.IntVar(&numProducers, "np", numProducers, "number of producers, np > 0")
	flag.IntVar(&numConsumers, "nc", numConsumers, "number of consumers, nc > 0")
	flag.Parse()
}
func main() {
	if maxMsgPerProducer < 1 || numProducers < 1 || numConsumers < 1 {
		flag.Usage()
		return
	}

	// make a channel big enough for the total possible messages
	ch = make(chan string)
	for i := 0; i < numProducers; i++ {
		producer(i + 1)
	}
	for i := 0; i < numConsumers; i++ {
		consumer(i + 1)
	}

	wgProducers.Wait()
	close(ch)
	wgConsumers.Wait()
}

// consumer processes the message sent to 'ch' by an unknown number of producers
func consumer(id int) {
	wgConsumers.Add(1)

	go func() {
		result := make(map[string]ProdInfo)
		for msg := range ch {
			fields := strings.Split(msg, ", num: ")
			v, _ := strconv.ParseInt(fields[1], 10, 64)
			// look up producer record, gets a copy
			pi := result[fields[0]]
			pi.sum += int(v)
			pi.count++
			result[fields[0]] = pi // store updated record
		}
		sum, count := 0, 0

		consumerMutex.Lock()
		fmt.Printf("Consumer %v\n", id)
		for k, v := range result {
			fmt.Printf("\t%v\n\t\tNumber of Elements: %v\n\t\tSub-total: %v\n", k, v.count, v.sum)
			sum += v.sum
			count += v.count
		}
		fmt.Printf("\tTotal count: %v\n", count)
		fmt.Printf("\tTotal sum: %v\n", sum)
		consumerMutex.Unlock()
		wgConsumers.Done()
	}()
}

// producer produces between 1 to maxMsgPerProducer string messages to 'ch'
// with the producer 'id' and a random int
func producer(id int) {
	wgProducers.Add(1)

	go func() {
		m := (r.Int() % maxMsgPerProducer) + 1
		for i := 0; i < m; i++ {
			ch <- fmt.Sprintf("Producer # %v, num: %v", id, r.Int()%20)
		}
		wgProducers.Done()
	}()
}
