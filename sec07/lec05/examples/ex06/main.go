// Section 07 - Lecture 05 : Channel Selection

package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// selecting from multiple channels
	// ----
	var ch1, ch2 chan string
	select {
	case <-ch1:
		log.Info("Got data *FROM* ch1")
	case ch2 <- "hello":
		log.Info("Sent data *TO* ch2")
	default:
		log.Warn("No communication for ch1 or ch2")
	}

}
