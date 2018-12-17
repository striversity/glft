package main

import (
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
)

type (
	cmdServer struct {
		conn net.Conn
	}
)

func main() {
	ln, err := net.Listen("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Infof("Command server listening %v", localhost)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}
func handleConnection(c net.Conn) {
	// Echo all incoming data.
	w := &cmdServer{conn: c}
	io.Copy(w, c)
	// Shut down the connection.
	c.Close()
}
func (wc *cmdServer) Write(p []byte) (n int, err error) {
	if wc == nil {
		return len(p), nil
	}
	io.WriteString(wc.conn, "[From server] ")
	return wc.conn.Write(p)
}
