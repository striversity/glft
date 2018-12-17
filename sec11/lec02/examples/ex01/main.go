// Section 11 - Lecture 02 : Simple HTTP Server v2
package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
)

var (
	rootRequests int
)

func main() {
	// serving dynamic data
	// ----
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootRequests++
	fmt.Fprintf(w, "This is rootHandler being called, %v times", rootRequests)
}
