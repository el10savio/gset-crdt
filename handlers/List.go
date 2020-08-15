package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// List ...
func List(w http.ResponseWriter, r *http.Request) {
	if len(GetPeerList()) != 0 {
		GSet, _ = Sync(GSet)
	}

	GSet := GSet.List()

	log.WithFields(log.Fields{
		"set": GSet,
	}).Debug("successful gset list")

	// json encode response value
	json.NewEncoder(w).Encode(GSet)
}
