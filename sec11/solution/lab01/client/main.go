package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/striversity/glft/sec11/solution/lab01/cmd"
)

type (
	writeCounter struct {
		count    int
		numBytes int
	}
)

var (
	host   = "localhost:12345"
	dryRun bool
)

func main() {
	flag.StringVar(&host, "s", host, "remote command server")
	flag.BoolVar(&dryRun, "n", false, "dry run, don't try to connect to server")
	flag.Parse()

	var conn net.Conn
	var err error
	if !dryRun {
		conn, err = net.Dial("tcp", host)
		if err != nil {
			log.Fatalf("Unable to connect to %v: %v", host, err)
		}
		defer conn.Close()
	}

	// remote cmd should be in flag.Args()[0], and the args in flag.Args()[1:]
	rcmd, err := encCmd(flag.Args())
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("Remote command: %v", rcmd)
	if dryRun {
		return
	}

	// send command to server
	_, err = io.WriteString(conn, rcmd)
	if err != nil {
		log.Errorf("Couldn't send command to server: %v", err)
		return
	}
	wg := new(sync.WaitGroup)
	setupCopy(wg, conn, os.Stdin)
	// TODO - have to remote output go to writeCounter for accounting
	setupCopy(wg, &writeCounter{}, conn)
	wg.Wait()
}
func setupCopy(wg *sync.WaitGroup, w io.Writer, r io.Reader) {
	wg.Add(1)
	go func() {
		io.Copy(w, r)
		wg.Done()
	}()
}

func encCmd(args []string) (rcmd string, err error) {
	if len(args) == 0 {
		err = fmt.Errorf("No commands provided")
		return
	}

	fmt.Printf("%v\n", args)
	if len(args) > 0 {
		rcmd = args[0]
		args = args[1:]
	}

	var jenc *json.Encoder
	var sbuf = new(strings.Builder)
	if jenc == nil {
		jenc = json.NewEncoder(sbuf)
	}
	v := &cmd.Command{Name: rcmd, Args: args}
	err = jenc.Encode(v)
	if err != nil {
		return
	}
	rcmd = sbuf.String()
	return
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	if wc == nil {
		return len(p), nil
	}
	wc.count++
	wc.numBytes += len(p)
	return os.Stdout.Write(p)
}
func (wc *writeCounter) String() string {
	if wc == nil {
		return ""
	}
	return fmt.Sprintf("\n[[Write Counter: <writes: %v, bytes: %v>]]\n", wc.count, wc.numBytes)
}
