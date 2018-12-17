// Section 11 - Lecture 04 : Generic TCP/IP Server & Client
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/striversity/glft/sec11/lec04/examples/ex03/proto"
)

var (
	remoteAddr = "localhost:12345"
)

func main() {
	flag.StringVar(&remoteAddr, "s", remoteAddr, "random number server address")
	flag.Parse()

	// return to basics - consumer using TCP/IP
	// ----
	consumer(remoteAddr)
}

// consumer is a client that consume messages from server
func consumer(server string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Errorf("Can't connect to server %v: %v", server, err)
		return
	}

	jdec := json.NewDecoder(conn)
	var msg proto.Message
	for {
		err = jdec.Decode(&msg)
		if err != nil {
			log.Errorf("Can't decode message from server %v: %v", server, err)
			return
		}
		processMessage(msg)
		conn.SetDeadline(time.Now().Add(5 * time.Second)) // reset timer after we receive work
	}
}

func processMessage(msg proto.Message) {
	total := 0

	for v := range msg.Payload {
		total += v
	}
	fmt.Printf("%v send:\n", msg.Source)
	l := len(msg.Payload)
	fmt.Printf("\tcount: %v, total: %v, avg: %v\n", l, total, float64(total)/float64(l))
}
