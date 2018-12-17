// Section 11 - Lecture 02 : Simple HTTP Server v2
package main

import (
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	localhost = ":12345"
)

func main() {
	// read request body
	// ----
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("No data from client\n")
		fmt.Fprintf(w, "No data provided\n")
		return
	}

	fmt.Printf("Client sent: '%v'\n", string(body))
	caps := strings.Title(string(body))
	fmt.Fprintf(w, "<p><b>Input:</b>%v</p>", string(body))
	fmt.Fprint(w, "<hr/>")
	fmt.Fprintf(w, "<p><b>Output:</b>%v</p>", caps)
}
