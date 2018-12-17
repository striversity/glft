// Section 11 - Lecture 02 : Simple HTTP Server v2
package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	localhost = ":12345"
)

var (
	rootRequests int
	fooRequests  int
	barRequests  int
)

func main() {
	// multiple handlers
	// ----
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fooRequests++
	fmt.Fprintf(w, "This is fooHandler being called, %v times", fooRequests)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	barRequests++
	fmt.Fprintf(w, "This is barHandler being called, %v times", barRequests)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootRequests++
	fmt.Fprintf(w, "This is rootHandler being called, %v times", rootRequests)
}
