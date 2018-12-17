// Section 11 - Lecture 04 : Generic TCP/IP Server & Client
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
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
	// return to basics - producer and consumer using TCP/IP
	// ----
	serverAddr := ":12345"
	remoteAddr := "localhost:12345"
	wg := new(sync.WaitGroup)
	producer(wg, "Producer 1", serverAddr)
	consumer(wg, remoteAddr)
	wg.Wait()
}

// consumer is a client that consume messages from server
func consumer(wg *sync.WaitGroup, server string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", server)
		if err != nil {
			log.Errorf("Can't connect to server %v: %v", server, err)
			return
		}

		jdec := json.NewDecoder(conn)
		var msg Message
		for {
			err = jdec.Decode(&msg)
			if err != nil {
				log.Errorf("Can't decode message from server %v: %v", server, err)
				return
			}
			processMessage(msg)
			conn.SetDeadline(time.Now().Add(5 * time.Second)) // reset timer after we receive work
		}
	}()
}

// producer is a server the provides messages for clients
func producer(wg *sync.WaitGroup, src string, addr string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Errorf("Can't create server for %v: %v", addr, err)
			return
		}
		defer ln.Close()
		log.Infof("Random number server started on %v", addr)

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Errorf("Client connection error: %v", err)
				continue // get another connection to serve
			}

			go serveClient(conn, src)
		}
	}()
}

func serveClient(conn net.Conn, src string) {
	jenc := json.NewEncoder(conn)
	var err error
	for {
		msg := newMessage(src)
		err = jenc.Encode(msg)
		if err != nil {
			log.Errorf("Can't send to client %v, closing connection: %v", conn.RemoteAddr(), err)
			conn.Close()
			return
		}
		time.Sleep(time.Duration(r.Intn(2000)) * time.Millisecond) // random delay up to 2sec
		conn.SetDeadline(time.Now().Add(5 * time.Second))          // reset timer after sending
	}
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
