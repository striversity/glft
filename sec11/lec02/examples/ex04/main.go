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

func main() {
	// getting request info
	// ----
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	remoteHost := r.RemoteAddr
	remoteURI := r.RequestURI
	method := r.Method
	fmt.Fprintf(w, "<p>Remote host/client addr: %v</p>", remoteHost)
	fmt.Fprintf(w, "<p>Remote host/client URI: %v</p>", remoteURI)
	fmt.Fprintf(w, "<p>Request made using HTTP.METHOD: %v</p>", method)
}
