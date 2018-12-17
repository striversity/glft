package main

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type (
	WriteCounter struct {
		count    int
		numBytes int
	}
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("No command specified")
	}

	name := os.Args[1]
	var cmdArgs []string
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}
	cmd := exec.Command(name, cmdArgs...)
	cmd.Stdin = os.Stdin
	wc := new(WriteCounter)
	cmd.Stdout = wc
	err := cmd.Run()
	if err != nil {
		log.Errorf("Command err: %v", err)
	}
	fmt.Println(wc)
}

func (wc *WriteCounter) Write(p []byte) (n int, err error) {
	if wc == nil {
		return len(p), nil
	}
	wc.count++
	wc.numBytes += len(p)
	return os.Stdout.Write(p)
}
func (wc *WriteCounter) String() string {
	if wc == nil {
		return ""
	}
	return fmt.Sprintf("\n[[Write Counter: <writes: %v, bytes: %v>]]\n", wc.count, wc.numBytes)
}
