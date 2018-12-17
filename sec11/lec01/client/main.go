// Section 11 - Lecture 01 : Simple HTTP Client & Server
package main

import (
	"io"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const (
	localhost = "localhost:12345"
	host      = "https://google.com"
)

func main() {
	// client - using the net/http package to performat GET request
	// ----
	resp, err := http.Get(host)
	if err != nil {
		log.Printf("Can't connect to %v: %v", host, err)
		log.Printf("Trying %v ...", localhost)
		resp, err = http.Get(localhost)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
