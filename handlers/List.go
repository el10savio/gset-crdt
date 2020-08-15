package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// List is the HTTP handler used to return
// all the values present in the GSet node in the server
func List(w http.ResponseWriter, r *http.Request) {
	// Sync the GSets if multiple nodes
	// are present in a cluster
	if len(GetPeerList()) != 0 {
		GSet, _ = Sync(GSet)
	}

	// Get the values from the GSet
	set := GSet.List()

	// DEBUG log in the case of success
	// indicating the new GSet
	log.WithFields(log.Fields{
		"set": set,
	}).Debug("successful gset list")

	// JSON encode response value
	json.NewEncoder(w).Encode(set)
}
