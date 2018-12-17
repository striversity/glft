// Section 07 - Lecture 05 : Channel Selection

package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// the 'select' statement
	// ----
	var ch1 chan int
	select {
	case <-ch1:
		log.Info("Got data from ch1")
	default:
		log.Warn("No data from ch1")
	}

}
