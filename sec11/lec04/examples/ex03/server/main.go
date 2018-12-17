// Section 11 - Lecture 04 : Generic TCP/IP Server & Client
package main

import (
	"encoding/json"
	"math/rand"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/sec11/lec04/examples/ex03/proto"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	// return to basics - producer and consumer using TCP/IP
	// ----
	serverAddr := ":12345"
	wg := new(sync.WaitGroup)
	producer(wg, "Producer 1", serverAddr)
	wg.Wait()
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

func newMessage(src string) proto.Message {
	count := r.Intn(10) + 1 // 1 to 10
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = r.Intn(100)
	}
	return proto.Message{Source: src, Payload: data}
}
