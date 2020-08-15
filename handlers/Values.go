package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Values ...
func Values(w http.ResponseWriter, r *http.Request) {
	GSet := GSet.List()

	log.WithFields(log.Fields{
		"set": GSet,
	}).Debug("successful gset list")

	// json encode response value
	json.NewEncoder(w).Encode(GSet)
}
