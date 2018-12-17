/* 
TODO 1 - Write a complete Go program which prints the 
'Featured articles' on the Go Language website: https://golang.org
TIP : Copy the text and print it using a raw string. */

package main

import "fmt"

func main() {
	fmt.Println(`
Featured articles
A Proposal for Package Versioning in Go
Eight years ago, the Go team introduced goinstall (which led to go get) and with it the decentralized, URL-like import paths that Go developers are familiar with today. After we released goinstall, one of the first questions people asked was how to incorporate version information. We admitted we didn’t know. For a long time, we believed that the problem of package versioning would be best solved by an add-on tool, and we encouraged people to create one. The Go community created many tools with different approaches. Each one helped us all better understand the problem, but by mid-2015 it was clear that there were now too many solutions. We needed to adopt a single, official tool.
Published 26 March 2018

Go 2017 Survey Results
This post summarizes the result of our 2017 user survey along with commentary and insights. It also draws key comparisons between the results of the 2016 and 2017 survey.

Published 26 February 2018
		`)
}
