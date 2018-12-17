package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os/exec"
	"time"

	"github.com/striversity/glft/sec11/solution/lab01/cmd"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
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
func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // wait 5 secs for command before closing connection

	// log.Infof("Waiting for command from client %v...", c.RemoteAddr())
	// read command from client
	jdec := json.NewDecoder(conn)
	remoteCmd := new(cmd.Command)
	err := jdec.Decode(remoteCmd)
	// log.Infof("Got  command from client...%v", remoteCmd)
	if err != nil {
		// try to send an message about the error and close connection if we can't read from client
		io.WriteString(conn, fmt.Sprintf("[REMOTE ERR] - %v", err))
		return
	}

	conn.SetDeadline(time.Now().Add(1 * time.Minute)) // close conn if idle for 5mins
	localCmd := exec.Command(remoteCmd.Name, remoteCmd.Args...)
	localCmd.Stdin = conn
	localCmd.Stdout = conn
	localCmd.Stderr = conn

	err = localCmd.Run()
	if err != nil {
		log.Errorf("Failed to run command %v for client %v", remoteCmd.Name, conn.RemoteAddr())
		io.WriteString(conn, fmt.Sprintf("[REMOTE CMD EXIT STATUS] - %v", err))
		return
	}
	// log.Infof("Command completed successfully for client %v", c.RemoteAddr())
}
