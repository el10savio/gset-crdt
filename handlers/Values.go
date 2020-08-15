package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Values ...
func Values(w http.ResponseWriter, r *http.Request) {
	set := GSet.List()

	log.WithFields(log.Fields{
		"set": set,
	}).Debug("successful gset values")

	// json encode response value
	json.NewEncoder(w).Encode(set)
}
