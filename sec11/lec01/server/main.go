// Section 11 - Lecture 01 : Simple HTTP Client & Server
package main

import (
	"fmt"
	"io"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	localhost = ":12345"
)

func main() {
	// server - using the net/http package to create an HTTP server
	// ----
	fmt.Printf("Server listening on %v\n", localhost)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(localhost, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const respText = `
		<html>
			<head>
				<title>Hello from the Web</title>
			</head>
			<body>
				<p>
					<b>Hello, World\n</b>
				</p>
				<p>
					My name is <i><b>Verrol</b></i>
				</p>
			</body>
		</html>
	`
	io.WriteString(w, respText)
}
