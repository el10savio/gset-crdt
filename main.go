package main

// The following implements the main Go
// package starting up the gset server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/gset-crdt/handlers"
)

const (
	// PORT ...
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started GSet node server")

	http.ListenAndServe(":"+PORT, r)
}
