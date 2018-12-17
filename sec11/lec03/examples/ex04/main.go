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
	countRequests int
)

func main() {
	// pitfall - registering multiple handlers for the same 'path'
	// ----
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", rootHandler)
	s1 := &http.Server{
		Addr:    localhost,
		Handler: new(ServerOne),
	}
	log.Fatal(s1.ListenAndServe())
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
	countRequests++
	fmt.Printf("rootHandler() called, %v times\n", countRequests)
	fmt.Fprintf(w, "This is handler being called, %v times", countRequests)
}
