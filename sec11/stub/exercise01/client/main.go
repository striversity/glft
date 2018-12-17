package main

import (
	"flag"
	"fmt"
	"os"
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

	// TODO 1 - complete the main() function to fetch data from the server specified by 'hostname'
}

func pageStats(buf []byte) {
	// TODO 2 - Fetch the page and determine how many types of HTML tags are used alone with the page size
	fmt.Printf("Stats fro Page from '%v' has:\n", hostname)
}
