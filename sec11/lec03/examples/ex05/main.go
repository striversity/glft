// Section 11 - Lecture 03 : Implementing htt.Handler
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	ServerOne struct {
		counter int
	}
	ServerTwo struct {
		prefix string // prefix to add to each reponse
	}
)

const (
	srvAddr1 = ":12345"
	srvAddr2 = ":12346"
)

func main() {
	// server - using the net/http package to create an HTTP server
	// ----
	s1 := &http.Server{
		Addr:        srvAddr1,
		Handler:     new(ServerOne),
		IdleTimeout: 5 * time.Second,
	}
	s2 := &http.Server{
		Addr:        srvAddr2,
		Handler:     &ServerTwo{prefix: "Logger"},
		IdleTimeout: 10 * time.Second,
	}

	wg := new(sync.WaitGroup)
	startServer(wg, s1)
	startServer(wg, s2)
	wg.Wait()
	log.Info("Server app exiting...")
}

func (s *ServerOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s == nil {
		return
	}
	s.counter++
	fmt.Fprintf(w, "This is ServerOne being called, %v times", int(s.counter))
}

func (s *ServerTwo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s == nil {
		return
	}
	fmt.Fprintf(w, "[%s] - This is ServerTwo being called at %v", s.prefix, time.Now())
}

func startServer(wg *sync.WaitGroup, server *http.Server) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Infof("Starting server on addr %v", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Warnf("Server on addr %v didn't start: %v", server.Addr, err)
			return
		}
	}()
}
