// Section 07 - Lecture 05 : Channel Selection

package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	// inserting delays the wrong way
	// ----
	fmt.Printf("Message 1 at: %v\n", time.Now())
	sleep(1 * time.Second)
	fmt.Printf("Message 2 at: %v\n", time.Now())

}
func sleep(delay time.Duration) {
	end := time.Now().Add(delay)
	for time.Now().Before(end) {
		log.Info("Just waiting here, nothing to do...")
	}
}
