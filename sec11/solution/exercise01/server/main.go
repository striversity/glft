package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
	respFile  = "./index.html"
)

var (
	buf *bytes.Buffer
)

func main() {
	// serve contents from the file 'index.html'
	// ----
	data, err := ioutil.ReadFile(respFile)
	if err != nil {
		log.Fatalf("Unable to read '%v': %v", respFile, err)
	}

	buf = bytes.NewBuffer(data)

	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, buf)
}
