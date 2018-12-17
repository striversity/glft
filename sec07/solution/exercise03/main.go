package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	maxMsgPerProducer = 100
	numProducers      = 3
	s                 = rand.NewSource(time.Now().Unix())
	r                 = rand.New(s)
	ch                chan string
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
	flag.Parse()
}
func main() {
	if maxMsgPerProducer < 1 || numProducers < 1 {
		flag.Usage()
		return
	}

	// make a channel big enough for the total possible messages
	ch = make(chan string, maxMsgPerProducer*numProducers)
	for i := 0; i < numProducers; i++ {
		producer(i + 1)
	}
	close(ch)
	consumer()
}

// consumer processes the message sent to 'ch' by an unknown number of producers
func consumer() {
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
	for k, v := range result {
		fmt.Printf("%v\n Number of Elements: %v\n Sub-total: %v\n", k, v.count, v.sum)
		sum += v.sum
		count += v.count
	}
	fmt.Printf("Total count: %v\n", count)
	fmt.Printf("Total sum: %v\n", sum)
}

// producer produces between 1 to maxMsgPerProducer string messages to 'ch'
// with the producer 'id' and a random int
func producer(id int) {
	m := (r.Int() % maxMsgPerProducer) + 1
	for i := 0; i < m; i++ {
		ch <- fmt.Sprintf("Producer # %v, num: %v", id, r.Int()%20)
	}
}
