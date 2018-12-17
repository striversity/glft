package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var hostname string
var help bool

func init() {
	flag.BoolVar(&help, "h", false, "show usage")
	flag.StringVar(&hostname, "s", "http://localhost:12345", "website to fetch, eg: https://google.com")
	flag.Parse()
}

func main() {
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// TODO 1 - TODO 1 - complete the main() function to fetch data from the server specified by 'hostname'
	resp, err := http.Get(hostname)
	if nil != err {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var buf bytes.Buffer // needs no initialization
	io.Copy(&buf, resp.Body)
	pageStats(buf.Bytes())
}

func pageStats(buf []byte) {
	// TODO 2 - Fetch the page and determine how many types of HTML tags are used alone with the page size
	fmt.Printf("Stats fro Page from '%v' has:\n", hostname)
	fmt.Printf("\tPage size: %v bytes\n", len(buf))
	fmt.Printf("\tImage (img-tags): %v\n", bytes.Count(buf, []byte("<img")))
	fmt.Printf("\tLinks (a-tags): %v\n", bytes.Count(buf, []byte("<a")))
	fmt.Printf("\tDivs (div-tags): %v\n", bytes.Count(buf, []byte("<div")))
	fmt.Printf("\tForms (form-tags): %v\n", bytes.Count(buf, []byte("<form")))
	fmt.Printf("\tButtons (button-tags): %v\n", bytes.Count(buf, []byte("<button")))
	fmt.Printf("\tInput (input-tags): %v\n", bytes.Count(buf, []byte("<input")))
}
