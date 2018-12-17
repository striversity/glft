// Section 11 - Lecture 03 : Implementing htt.Handler
package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
)

type (
	// ServerOne implements http.Handler
	ServerOne int
)

var (
	rootRequests int
)

func main() {
	// creating an HTTP Server by implementing net/http.Handler
	// ----
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", rootHandler)
	http.Handle("/bar", new(ServerOne))
	log.Fatal(http.ListenAndServe(localhost, nil))
}

func (s *ServerOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s == nil {
		return
	}
	*s++
	fmt.Printf("ServerOne.ServeHTTP() called, %v times\n", *s)
	fmt.Fprintf(w, "This is ServerOne being called, %v times", *s)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootRequests++
	fmt.Fprintf(w, "This is rootHandler being called, %v times", rootRequests)
}
