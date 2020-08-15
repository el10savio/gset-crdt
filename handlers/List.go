package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// List ...
func List(w http.ResponseWriter, r *http.Request) {
	set := GSet.List()

	log.WithFields(log.Fields{
		"set": set,
	}).Debug("successful gset list")

	// json encode response value
	json.NewEncoder(w).Encode(set)
}
